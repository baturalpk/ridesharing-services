package common

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"

	pb "github.com/baturalpk/ridesharing-services/client-app/genproto"
)

func AskLocation(q string) (*pb.Location, error) {
	ans := struct {
		Lat  float32
		Long float32
	}{}
	fmt.Println(q)
	err := survey.Ask([]*survey.Question{
		{
			Name:   "lat",
			Prompt: &survey.Input{Message: "Latitude: "},
		},
		{
			Name:   "long",
			Prompt: &survey.Input{Message: "Longitude: "},
		},
	}, &ans)
	if err != nil {
		return nil, err
	}
	return &pb.Location{
		Latitude:  ans.Lat,
		Longitude: ans.Long,
	}, nil
}
