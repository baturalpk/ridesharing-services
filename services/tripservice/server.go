package main

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/baturalpk/ridesharing-services/services/tripservice/genproto"
	"github.com/baturalpk/ridesharing-services/services/tripservice/trip"
)

const (
	TripSeekingRadius = 0.5
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
		RiderID: req.GetRiderId(),
		Start: trip.Location{
			Lat:  req.GetStart().GetLatitude(),
			Long: req.GetStart().GetLongitude(),
		},
		Destination: trip.Location{
			Lat:  req.GetDestination().GetLatitude(),
			Long: req.GetDestination().GetLongitude(),
		},
		RideType: req.GetRideType(),
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
		return nil, status.Errorf(codes.Internal, err.Reason)
	}
	return &proto.Empty{}, nil
}

func (s *server) SeekForTripAsDriver(driverSv proto.TripService_SeekForTripAsDriverServer) error {
	for {
		req, err := driverSv.Recv()
		if err != nil {
			return status.Error(codes.Internal, "Internal error")
		}

		dloc := req.GetDriverLocation()
		ts, err2 := s.repo.GetByStartLocationInterval(
			driverSv.Context(),
			trip.Location{
				Lat:  dloc.Latitude - TripSeekingRadius,
				Long: dloc.Longitude - TripSeekingRadius,
			},
			trip.Location{
				Lat:  dloc.Latitude + TripSeekingRadius,
				Long: dloc.Longitude + TripSeekingRadius,
			},
		)
		if len(ts) > 0 && err2 == nil {
			t := ts[0]
			return driverSv.Send(&proto.SeekForTripAsDriverResponse{
				TripId:  t.ID.Hex(),
				RiderId: t.RiderID,
				Start: &proto.Location{
					Latitude:  t.Start.Lat,
					Longitude: t.Start.Long,
				},
				Destination: &proto.Location{
					Latitude:  t.Destination.Lat,
					Longitude: t.Destination.Long,
				},
			})
		}
	}
}
