package bootstrap

import (
	"fmt"
	"time"

	aggreation "github.com/lagrange92/Haechi/aggregation"
	"github.com/lagrange92/Haechi/model"
	"github.com/lagrange92/Haechi/singleton"
)

// Bootstrap : load seoul spots data and request latest ppl data
func Bootstrap() {
	bootstrapSeoulSpot()
	bootstrapLatestPpl(singleton.SeoulSpots)
}

// ActivateWorker : run update ppl worker eternally, should invoke as goroutine
func ActivateWorker() {
	time.Sleep(1 * time.Minute) // wait 1 minute for bootstrap finished at first

	activateUpdatePplWorker()
}

func activateUpdatePplWorker() {
	updatePplWorker := time.NewTicker(5*time.Minute + 15*time.Second)
	defer updatePplWorker.Stop()

	for {
		select {
		case <-updatePplWorker.C:
			fmt.Println("Updating ppl data started...")

			bootstrapLatestPpl(singleton.SeoulSpots)

			fmt.Println("Updating ppl data finished.")
		}
	}
}

func bootstrapSeoulSpot() {
	seoulSpots := loadSeoulSpots()
	singleton.SeoulSpots = seoulSpots
}

func bootstrapLatestPpl(seoulSpots []model.SeoulSpot) {
	latestPpl := aggreation.AggregatePpl(seoulSpots)
	singleton.LatestPpl = latestPpl
}
