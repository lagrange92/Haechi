package bootstrap

import (
	aggreation "github.com/lagrange92/Haechi/aggregation"
	"github.com/lagrange92/Haechi/singleton"
)

// Load : load seoul spots data and latest ppl data
func Load() {
	seoulSpots := loadSeoulSpots()

	latestPpl := aggreation.AggregatePpl(seoulSpots)

	singleton.SeoulSpots = seoulSpots
	singleton.LatestPpl = latestPpl
}
