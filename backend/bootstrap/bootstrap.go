package bootstrap

import (
	"fmt"
	"time"

	aggreation "github.com/lagrange92/Haechi/aggregation"
	"github.com/lagrange92/Haechi/model"
	"github.com/lagrange92/Haechi/store"
)

// Bootstrap : load seoul spots data and request latest ppl data
func Bootstrap() {
	bootstrapSeoulSpot()
	bootstrapLatestPpl(store.SeoulSpots)
}

// ActivateWorker : run update ppl worker eternally, should invoke as goroutine
func ActivateWorker() {
	time.Sleep(1 * time.Hour) // test code

	activateUpdatePplWorker()
}

func activateUpdatePplWorker() {
	updatePplWorker := time.NewTicker(5*time.Minute + 15*time.Second)
	defer updatePplWorker.Stop()

	for {
		select {
		case <-updatePplWorker.C:
			fmt.Println("Updating ppl data started...")

			bootstrapLatestPpl(store.SeoulSpots)

			fmt.Println("Updating ppl data finished.")
		}
	}
}

func bootstrapSeoulSpot() {
	seoulSpots := loadSeoulSpots()
	store.SeoulSpots = seoulSpots
}

func bootstrapLatestPpl(seoulSpots []model.SeoulSpot) {
	latestPpl := aggreation.AggregatePpl(seoulSpots)
	store.PplDistribution = latestPpl
}
