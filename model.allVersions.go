package main

import (
	"encoding/json"
)

//AllVersions to contain the version number array
type AllVersions struct {
	Versions []string
}

//ToAllVersionsJSON to be used for marshalling of CI type
func (b AllVersions) ToAllVersionsJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

//FromAllVersionsJSON to be used for unmarshalling of CI type
func FromAllVersionsJSON(data []byte) []string {
	book := make([]string, 0)
	err := json.Unmarshal(data, &book)
    //err := mapstructure.Decode(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}