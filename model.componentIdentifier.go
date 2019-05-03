package main

import (
	"encoding/json"
)

//CI to contain the component identifier
type CI struct {
	ComponentIdentifier ComponentIdentifier  `json:"componentIdentifier"`
}

//ComponentIdentifier to contain the format and coordinate JSON
type ComponentIdentifier struct {
	Format        string `json:"format"`
	Coordinates   MavenCoordinate `json:"coordinates"`
}

//MavenCoordinate to contain the GAVTC
type MavenCoordinate struct {
	ArtifactID string `json:"artifactId"`
	Classifier string `json:"classifier"`
	Extension string `json:"extension"`
	GroupID string `json:"groupId"`
	Version string `json:"version"`
}

//ToComponentIdentifierJSON to be used for marshalling of CI type
func (b CI) ToComponentIdentifierJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

//ToComponentIdentifierJSON to be used for marshalling of CI type
func (b ComponentIdentifier) ToComponentIdentifierJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

//FromComponentIdentifierJSON to be used for unmarshalling of CI type
func FromComponentIdentifierJSON(data []byte) CI {
	book := CI{}
	err := json.Unmarshal(data, &book)
    //err := mapstructure.Decode(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}