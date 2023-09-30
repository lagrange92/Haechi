package utils

import (
	"encoding/json"
	"fmt"

	"github.com/lagrange92/Haechi/model"
)

// Unmarshal : unmarshal json to model.SeoulCityJSON
func Unmarshal(from []byte) (to model.SeoulCityJSON) {
	errMarshal := json.Unmarshal(from, &to)
	if errMarshal != nil {
		fmt.Println(errMarshal)
	}

	return
}
