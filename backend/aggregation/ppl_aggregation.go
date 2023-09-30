package aggreation

import (
	"github.com/lagrange92/Haechi/model"
)

// AggregatePpl aggregates population data from Seoul Open API
func AggregatePpl(seoulSpots []model.SeoulSpot) []model.PpltnData {
	ppls := []model.PpltnData{}
	pplCh := make(chan model.PpltnData)

	// Request latest ppl data to Seoul Open API by using goroutines
	for _, seoulSpot := range seoulSpots {
		go requestPpl(seoulSpot, pplCh)
	}

	// Receive ppl data from channel
	for i := 0; i < len(seoulSpots); i++ {
		chanData := <-pplCh

		// abandon empty data
		if chanData.AreaName == "" {
			continue
		}

		ppls = append(ppls, chanData)
	}

	normalizePpl(ppls)

	return ppls
}
