package trip

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
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
	c *mongo.Collection
}

func NewRepository(coll *mongo.Collection) *Repository {
	return &Repository{c: coll}
}

func (r *Repository) CreateTrip(ctx context.Context, trip Trip) (string, *TripError) {
	trip.CreatedAt = primitive.Timestamp{T: uint32(time.Now().Unix())}
	doc, err := r.c.InsertOne(ctx, trip)
	if err != nil {
		return "", NewInternalTripError(err)
	}
	id := doc.InsertedID.(primitive.ObjectID)
	return id.Hex(), nil
}

func (r *Repository) DeleteTrip(ctx context.Context, tripID string) *TripError {
	id, err := primitive.ObjectIDFromHex(tripID)
	if err != nil {
		return NewClientRelatedTripError(err, "the provided 'trip_id' is not valid")
	}
	if _, err = r.c.DeleteOne(ctx, bson.M{"_id": id}); err != nil {
		return NewInternalTripError(err)
	}
	return nil
}
