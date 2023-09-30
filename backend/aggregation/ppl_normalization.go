package aggreation

import (
	"fmt"
	"strconv"

	"github.com/lagrange92/Haechi/model"
)

// normalizePpl normalizes average population to 0~1
func normalizePpl(ppls []model.PplData) {
	normalMin := 987654321
	normalMax := 0

	for _, ppl := range ppls {
		avgPpltn, _ := strconv.Atoi(ppl.AreaAvgPpltn)

		if avgPpltn < normalMin {
			normalMin = avgPpltn
		} else if avgPpltn > normalMax {
			normalMax = avgPpltn
		}
	}

	for i, ppl := range ppls {
		avgPpltn, _ := strconv.Atoi(ppl.AreaAvgPpltn)
		normalPpltn := float64(avgPpltn-normalMin) / float64(normalMax-normalMin)

		ppls[i].AreaAvgPpltn = fmt.Sprintf("%f", normalPpltn)
	}
}
