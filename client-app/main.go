package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/baturalpk/ridesharing-services/client-app/common"
	pb "github.com/baturalpk/ridesharing-services/client-app/genproto"
	"github.com/baturalpk/ridesharing-services/client-app/mapping"
	"github.com/baturalpk/ridesharing-services/client-app/pricing"
	"github.com/baturalpk/ridesharing-services/client-app/trip"
)

const (
	TripSvAddr    = "127.0.0.1:50001"
	MappingSvAddr = "127.0.0.1:50002"
	PricingSvAddr = "127.0.0.1:50003"
)

var (
	qMain = []string{
		"Trip service client",
		"Mapping service client",
		"Pricing service client",
	}
)

func main() {
	fmt.Println("====================================")
	fmt.Println("🚖 Welcome to ridesharing-services 👋")
	fmt.Println("====================================")
	fmt.Println(common.Info("🚶 Press [Control + C] to leave anytime..."))

	var ans string
	err := survey.AskOne(&survey.Select{
		Message: "Which type of client would you like to create?",
		Options: qMain,
		Default: qMain[0],
	}, &ans)
	if err != nil {
		common.EvalError(err)
		return
	}

	switch ans {
	case qMain[0]:
		trip.Handler(pb.NewTripServiceClient(newConnection(TripSvAddr)))
	case qMain[1]:
		mapping.Handler(pb.NewMappingServiceClient(newConnection(MappingSvAddr)))
	case qMain[2]:
		pricing.Handler(pb.NewPricingServiceClient(newConnection(PricingSvAddr)))
	}
}

func newConnection(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		common.EvalError(err)
	}
	return conn
}
