package trip

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Trip struct {
	RiderID     interface{} `bson:"rider_id"`
	Start       interface{} `bson:"start"`
	Destination interface{} `bson:"destination"`
	RideType    interface{} `bson:"ride_type"`
}

type ErrorType int

const (
	InternalError ErrorType = iota
	ClientRelatedError
)

type TripError struct {
	Err    error
	Type   ErrorType
	Reason string
}

type Repository struct {
	c *mongo.Collection
}

func NewRepository(coll *mongo.Collection) *Repository {
	return &Repository{c: coll}
}

func (r *Repository) CreateTrip(ctx context.Context, trip Trip) (string, *TripError) {
	doc, err := r.c.InsertOne(ctx, trip)
	if err != nil {
		return "", &TripError{
			Err:    err,
			Type:   InternalError,
			Reason: "Internal error",
		}
	}
	id := doc.InsertedID.(primitive.ObjectID)
	return id.Hex(), nil
}

func (r *Repository) DeleteTrip(ctx context.Context, tripID string) *TripError {
	id, err := primitive.ObjectIDFromHex(tripID)
	if err != nil {
		return &TripError{
			Err:    err,
			Type:   ClientRelatedError,
			Reason: "the provided 'trip_id' is not valid",
		}
	}
	if _, err = r.c.DeleteOne(ctx, bson.M{"_id": id}); err != nil {
		return &TripError{
			Err:    err,
			Type:   InternalError,
			Reason: "Internal error",
		}
	}
	return nil
}
