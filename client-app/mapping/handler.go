package mapping

import (
	"context"
	"fmt"
	"github.com/AlecAivazis/survey/v2"

	"github.com/baturalpk/ridesharing-services/client-app/common"
	pb "github.com/baturalpk/ridesharing-services/client-app/genproto"
)

var (
	qMain = []string{
		"Calculate total distance for your route",
		"Make a time estimation for your route",
		"Convert an address to its corresponding geolocation (Geocoding)",
	}
)

func askRoute() (*pb.Route, error) {
	fmt.Println("Route:")
	sl, err := common.AskLocation("Start location:")
	if err != nil {
		return nil, err
	}
	dl, err := common.AskLocation("Destination location:")
	if err != nil {
		return nil, err
	}
	return &pb.Route{
		Start:       sl,
		Destination: dl,
	}, nil
}

func Handler(c pb.MappingServiceClient) {
	for {
		var ans string
		err := survey.AskOne(&survey.Select{
			Message: "Please choose an action:",
			Options: qMain,
			Default: qMain[0],
		}, &ans)
		if err != nil {
			common.EvalError(err)
			continue
		}

		switch ans {
		case qMain[0]:
			r, err := askRoute()
			if err != nil {
				common.EvalError(err)
				continue
			}
			req := &pb.CalculateDistanceRequest{Route: r}
			res, err := c.CalculateDistance(context.Background(), req)
			if err != nil {
				common.EvalError(err)
				continue
			}
			common.PrintResult(res)

		case qMain[1]:
			r, err := askRoute()
			if err != nil {
				common.EvalError(err)
				continue
			}
			req := &pb.MakeTimeEstimationRequest{Route: r}
			res, err := c.MakeTimeEstimation(context.Background(), req)
			if err != nil {
				common.EvalError(err)
				continue
			}
			common.PrintResult(res)

		case qMain[2]:
			var ans string
			err := survey.AskOne(&survey.Input{Message: "Address:"}, &ans)
			if err != nil {
				common.EvalError(err)
				continue
			}
			req := &pb.DoGeocodingRequest{Address: ans}
			res, err := c.DoGeocoding(context.Background(), req)
			if err != nil {
				common.EvalError(err)
				continue
			}
			common.PrintResult(res)
		}
	}
}
