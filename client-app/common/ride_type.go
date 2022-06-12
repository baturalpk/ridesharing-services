package common

import (
	"github.com/AlecAivazis/survey/v2"

	pb "github.com/baturalpk/ridesharing-services/client-app/genproto"
)

var (
	q = []string{"Default", "Comfort", "Premium"}
)

func AskRideType(m string) (pb.RideType, error) {
	var ans string
	err := survey.AskOne(&survey.Select{
		Message: m,
		Options: q,
		Default: q[0],
	}, &ans)
	if err != nil {
		return 0, err
	}

	for i, s := range q {
		if ans == s {
			return pb.RideType(int32(i)), nil
		}
	}
	return 0, nil
}
