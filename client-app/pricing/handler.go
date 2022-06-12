package pricing

import (
	"context"
	"github.com/AlecAivazis/survey/v2"

	"github.com/baturalpk/ridesharing-services/client-app/common"
	pb "github.com/baturalpk/ridesharing-services/client-app/genproto"
)

var (
	q1 = []string{"Calculate total cost"}

	q = []*survey.Question{
		{
			Name: "Main",
			Prompt: &survey.Select{
				Message: "Please choose an action:",
				Options: q1,
				Default: q1[0],
			},
		},
		{
			Name:   "Distance",
			Prompt: &survey.Input{Message: "How far did you travel?"},
		},
	}

	a = struct {
		Main     string
		Distance float32
	}{}
)

func Handler(c pb.PricingServiceClient) {
	for {
		err := survey.Ask(q, &a)
		if err != nil {
			common.EvalError(err)
			continue
		}

		switch a.Main {
		case q1[0]:
			typ, err := common.AskRideType("Which service did you use in your trip?")
			if err != nil {
				common.EvalError(err)
				continue
			}
			req := &pb.CalculateCostRequest{
				Type:     typ,
				Distance: a.Distance,
			}
			res, err := c.CalculateCost(context.Background(), req)
			if err != nil {
				common.EvalError(err)
				continue
			}
			common.PrintResult(res)
		}
	}
}
