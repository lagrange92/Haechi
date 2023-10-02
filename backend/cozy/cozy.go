package cozy

import (
	"errors"

	"github.com/lagrange92/Haechi/model"
)

const cozyPlacesNum = 3

// GetCozyPlaces : update cozy places data
func GetCozyPlaces(pplDist []model.PplData) ([]model.CozyPlacesData, error) {
	if pplDist == nil {
		return []model.CozyPlacesData{}, errors.New("PplDistribution is nil")
	}

	cozyPlaces := []model.CozyPlacesData{}

	for idx, ppl := range pplDist {
		if idx == cozyPlacesNum {
			break
		}

		cozyPlaces = append(cozyPlaces, model.CozyPlacesData{
			AreaName:      ppl.AreaName,
			AreaLatitude:  ppl.AreaLatitude,
			AreaLongitude: ppl.AreaLongitude,
			AreaAvgPpltn:  ppl.AreaAvgPpltn,
		})
	}

	return cozyPlaces, nil
}
