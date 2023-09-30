package utils

import (
	"fmt"
	"strconv"

	"github.com/lagrange92/Haechi/model"
)

// Convert : convert model.SeoulCityJSON to model.PpltnData
func Convert(from model.SeoulCityJSON) (to model.PpltnData) {
	ppltnMin, _ := strconv.Atoi(from.Ppltn[0].AreaPpltnMin)
	ppltnMax, _ := strconv.Atoi(from.Ppltn[0].AreaPpltnMax)

	to.AreaName = from.Ppltn[0].AreaName
	to.AreaCode = from.Ppltn[0].AreaCode
	to.AreaAvgPpltn = fmt.Sprintf("%d", (ppltnMin+ppltnMax)/2)

	return
}
