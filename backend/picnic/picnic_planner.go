package picnic

import (
	"fmt"

	"github.com/lagrange92/Haechi/model"
	"github.com/lagrange92/Haechi/naver"
)

// MakePlan makes picnic plan
func MakePlan(req model.PicnicRequest) (model.PicnicResponse, error) {
	// TODO: Make picnic plan
	// 1. Read festival data from URL (Seoul)
	// 2. Read weather forcast data from URL (Seoul)
	// 3. Read traffic forcast data from URL (Seoul)
	// 4. Read population forcast data from URL (Seoul)
	// 5. Read route data from URL (Naver)
	nvDirResponse, _ := naver.ReqDir(req.Start, req.Goal)

	// for test
	fmt.Println(nvDirResponse)

	// 6. Make picnic plan

	resp := model.PicnicResponse{
		Start:    req.Start,
		Goal:     req.Goal,
		LayOvers: req.LayOvers,
	}

	return resp, nil
}
