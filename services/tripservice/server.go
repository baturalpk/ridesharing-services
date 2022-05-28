package main

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/baturalpk/ridesharing-services/services/tripservice/genproto"
	"github.com/baturalpk/ridesharing-services/services/tripservice/trip"
)

func RegisterServer(gsv *grpc.Server, r *trip.Repository) {
	sv := &server{repo: r}
	proto.RegisterTripServiceServer(gsv, sv)
}

type server struct {
	repo *trip.Repository
	proto.UnimplementedTripServiceServer
}

func (s *server) RequestTrip(ctx context.Context, req *proto.RequestTripRequest) (*proto.RequestTripResponse, error) {
	t := trip.Trip{
		RiderID:     req.GetRiderId(),
		Start:       req.GetStart(),
		Destination: req.GetDestination(),
		RideType:    req.GetRideType(),
	}
	id, err := s.repo.CreateTrip(ctx, t)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Reason)
	}
	return &proto.RequestTripResponse{TripId: id}, nil
}

func (s *server) CancelTrip(ctx context.Context, req *proto.CancelTripRequest) (*proto.Empty, error) {
	err := s.repo.DeleteTrip(ctx, req.GetTripId())
	if err != nil {
		if err.Type == trip.ClientRelatedError {
			return nil, status.Errorf(codes.InvalidArgument, err.Reason)
		}
		return nil, status.Errorf(codes.InvalidArgument, err.Reason)
	}
	return &proto.Empty{}, nil
}

func (s *server) SeekForTripAsDriver(driverSv proto.TripService_SeekForTripAsDriverServer) error {
	//TODO implement me
	panic("implement me")
}
