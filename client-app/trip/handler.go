package trip

import (
	"context"
	"log"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/google/uuid"

	"github.com/baturalpk/ridesharing-services/client-app/common"
	pb "github.com/baturalpk/ridesharing-services/client-app/genproto"
)

var (
	riderID = uuid.NewString()

	qMain = []string{
		"Rider 🚶",
		"Driver 🚖",
	}

	qRider = []string{
		"Request new trip",
		"Cancel the trip",
	}

	qDriver = []string{
		"Find someone who requests a trip near here",
	}
)

func Handler(c pb.TripServiceClient) {
	var ans string
	err := survey.AskOne(&survey.Select{
		Message: "Do you want to be a rider or a driver?",
		Options: qMain,
		Default: qMain[0],
	}, &ans)
	if err != nil {
		common.EvalError(err)
	}

	switch ans {
	case qMain[0]:
		riderHandler(c)
	case qMain[1]:
		driverHandler(c)
	}
}

func riderHandler(c pb.TripServiceClient) {
	for {
		var ans string
		err := survey.AskOne(&survey.Select{
			Message: "Please choose an action:",
			Options: qRider,
			Default: qRider[0],
		}, &ans)
		if err != nil {
			common.EvalError(err)
			continue
		}

		switch ans {
		case qRider[0]:
			sl, err := common.AskLocation("Start location:")
			if err != nil {
				common.EvalError(err)
				continue
			}
			dl, err := common.AskLocation("Destination location:")
			if err != nil {
				common.EvalError(err)
				continue
			}
			typ, err := common.AskRideType("Which service would you like to use in your trip?")
			if err != nil {
				common.EvalError(err)
				continue
			}
			req := &pb.RequestTripRequest{
				RiderId:     riderID,
				Start:       sl,
				Destination: dl,
				RideType:    typ,
			}
			res, err := c.RequestTrip(context.Background(), req)
			if err != nil {
				common.EvalError(err)
				continue
			}
			common.PrintResult(res)

		case qRider[1]:
			var ans string
			err := survey.AskOne(&survey.Input{Message: "Which trip would you like to cancel? (type its id)"}, &ans)
			if err != nil {
				common.EvalError(err)
				continue
			}
			_, err = c.CancelTrip(context.Background(), &pb.CancelTripRequest{TripId: ans})
			if err != nil {
				common.EvalError(err)
				continue
			}
			common.PrintResult("success")
		}
	}
}

type driverChanResponse struct {
	res interface{}
	err error
}

func driverHandler(c pb.TripServiceClient) {
	for {
		var ans string
		err := survey.AskOne(&survey.Select{
			Message: "Please choose an action:",
			Options: qDriver,
			Default: qDriver[0],
		}, &ans)
		if err != nil {
			common.EvalError(err)
			continue
		}

		switch ans {
		case qDriver[0]:
			loc, err := common.AskLocation("Your current location:")
			if err != nil {
				common.EvalError(err)
				continue
			}

			stream, err := c.SeekForTripAsDriver(context.Background())
			if err != nil {
				common.EvalError(err)
				continue
			}

			ch := make(chan driverChanResponse, 1)
			done := make(chan struct{})
			go driverStreamSender(done, stream, loc)
			go driverStreamReceiver(ch, done, stream)
			log.Println(common.Info("Searching..."))
			ret := <-ch
			if ret.err != nil {
				common.EvalError(ret.err)
				continue
			}
			common.PrintResult(ret.res)
		}
	}
}

func driverStreamSender(done chan struct{}, stream pb.TripService_SeekForTripAsDriverClient, initialLoc *pb.Location) {
	for {
		select {
		case <-done:
			return
		default:
			if err := stream.Send(&pb.SeekForTripAsDriverRequest{DriverLocation: initialLoc}); err != nil {
				common.EvalError(err)
				return
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func driverStreamReceiver(ch chan driverChanResponse, done chan struct{}, stream pb.TripService_SeekForTripAsDriverClient) {
	res, err := stream.Recv()
	_ = stream.CloseSend()
	if err != nil {
		ch <- driverChanResponse{err: err}
		done <- struct{}{}
		return
	}
	ch <- driverChanResponse{res: res}
	done <- struct{}{}
}
