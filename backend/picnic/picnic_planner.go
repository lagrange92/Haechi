package picnic

import (
	"github.com/lagrange92/Haechi/model"
)

// MakePlan makes picnic plan
func MakePlan(req model.PicnicRequest) (model.PicnicResponse, error) {
	resp := model.PicnicResponse{
		Start:    req.Start,
		Goal:     req.Goal,
		LayOvers: req.LayOvers,
	}

	return resp, nil
}
