package aggreation

import (
	"sort"

	"github.com/lagrange92/Haechi/model"
)

// AggregatePpl aggregates population data from Seoul Open API
func AggregatePpl(seoulSpots []model.SeoulSpot) []model.PplData {
	ppls := []model.PplData{}
	pplCh := make(chan model.PplData)

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

	// sort by area average population
	sort.Slice(ppls, func(i, j int) bool {
		return ppls[i].AreaAvgPpltn < ppls[j].AreaAvgPpltn
	})

	return ppls
}
