package trip

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/go-redis/redis/v9"
)

const (
	GeoStoreKey = "trip-geo"
)

type Trip struct {
	ID          primitive.ObjectID  `bson:"_id,omitempty"`
	RiderID     string              `bson:"rider_id"`
	Start       Location            `bson:"start"`
	Destination Location            `bson:"destination"`
	RideType    interface{}         `bson:"ride_type"`
	CreatedAt   primitive.Timestamp `bson:"created_at"`
}

type Location struct {
	Lat, Long float32
}

type Repository struct {
	persist *mongo.Collection
	geodb   *redis.Client
}

func NewRepository(coll *mongo.Collection, rcl *redis.Client) *Repository {
	return &Repository{
		persist: coll,
		geodb:   rcl,
	}
}

func (r *Repository) CreateTrip(ctx context.Context, trip Trip) (string, *TripError) {
	trip.CreatedAt = primitive.Timestamp{T: uint32(time.Now().Unix())}
	doc, err := r.persist.InsertOne(ctx, trip)
	if err != nil {
		return "", NewInternalTripError(err)
	}
	trip.ID = doc.InsertedID.(primitive.ObjectID)

	if err = r.saveStartLocation(ctx, trip); err != nil {
		return "", NewInternalTripError(err)
	}
	return trip.ID.Hex(), nil
}

func (r *Repository) DeleteTrip(ctx context.Context, tripID string) *TripError {
	id, err := primitive.ObjectIDFromHex(tripID)
	if err != nil {
		return NewClientRelatedTripError(err, "the provided 'trip_id' is not valid")
	}
	if _, err = r.persist.DeleteOne(ctx, bson.M{"_id": id}); err != nil {
		return NewInternalTripError(err)
	}
	return nil
}

func (r *Repository) saveStartLocation(ctx context.Context, t Trip) error {
	cmd := r.geodb.Do(ctx, "GEOADD", GeoStoreKey, float64(t.Start.Long), float64(t.Start.Lat), t.ID.Hex())
	return cmd.Err()
}

func (r *Repository) SearchTripByRadiusMeter(ctx context.Context, rad float64, loc Location) (*Trip, *TripError) {
	// Search a proper trip by radius
	cmd := r.geodb.Do(ctx, "GEORADIUS", GeoStoreKey, float64(loc.Long), float64(loc.Lat), rad, "M")
	val, err := cmd.StringSlice()
	if err != nil {
		return nil, NewInternalTripError(err)
	}
	if len(val) <= 0 {
		return nil, nil
	}
	// If found, ...
	// ...choose the first candidate
	tid := val[0]
	oid, _ := primitive.ObjectIDFromHex(tid)

	// ...remove from geo store
	r.geodb.ZRem(ctx, GeoStoreKey, tid)

	// ...fetch full details about the trip
	doc := r.persist.FindOne(ctx, bson.M{"_id": oid})
	if err = doc.Err(); err != nil {
		return nil, NewInternalTripError(err)
	}

	var t Trip
	if err = doc.Decode(&t); err != nil {
		return nil, NewInternalTripError(err)
	}
	return &t, nil
}
