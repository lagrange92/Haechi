package utils

import (
	"encoding/json"
	"fmt"

	"github.com/lagrange92/Haechi/model"
)

// UnmarshalSeoulCity : unmarshal json to model.SeoulCityJSON
func UnmarshalSeoulCity(from []byte) (to model.SeoulCityJSON) {
	err := json.Unmarshal(from, &to)
	if err != nil {
		fmt.Println(err)
	}

	return
}

// UnmarshalKaoKwd : unmarshal json to model.KaoKwdResponse
func UnmarshalKaoKwd(from []byte) (to model.KaoKwdResponse) {
	err := json.Unmarshal(from, &to)
	if err != nil {
		fmt.Println(err)
	}

	return
}

// UnmarshalNVDir : unmarshal json to model.NVDirResponse
func UnmarshalNVDir(from []byte) (to model.NVDirResponse) {
	err := json.Unmarshal(from, &to)
	if err != nil {
		fmt.Println(err)
	}

	return
}
