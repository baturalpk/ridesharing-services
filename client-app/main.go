package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/baturalpk/ridesharing-services/client-app/genproto"
)

func main() {
	address := "127.0.0.1:8080"

	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("grpc: failed to dial: %v", err)
	}

	tripclient := pb.NewTripServiceClient(conn)

	resp1, err := tripclient.RequestTrip(context.Background(), &pb.RequestTripRequest{
		RiderId: "1234",
		Start: &pb.Location{
			Latitude:  3,
			Longitude: 4,
		},
		Destination: &pb.Location{
			Latitude:  5,
			Longitude: 12,
		},
		RideType: pb.RideType_COMFORT,
	})

	if err != nil {
		log.Fatalf("trip: requestTrip: error: %v", err)
	}
	tid := resp1.GetTripId()
	log.Printf("Requested a new trip, id: %s", tid)

	log.Println("the trip request will be canceled in 10 seconds...")
	time.Sleep(10 * time.Second)

	_, err = tripclient.CancelTrip(context.Background(), &pb.CancelTripRequest{
		TripId: tid,
	})
	if err != nil {
		log.Fatalf("trip: cancelTrip: error: %v", err)
	}
	log.Printf("Cancelled the trip (id: %s)", tid)
}
