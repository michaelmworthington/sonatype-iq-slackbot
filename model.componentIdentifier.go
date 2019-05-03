package main

import (
	"errors"
	"encoding/json"
	"strings"
	"regexp"
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

//ParseMavenCoordinate read a string and create a MavenCoordinate
func ParseMavenCoordinate(pText string) (MavenCoordinate, error) {
	var gav MavenCoordinate
	
	var parts = strings.Split(pText, ":")

	if (len(parts) != 3) {
		return gav, errors.New("Could not parse coordinate. Not enough parts: " + pText)
	}

	reg, err := regexp.Compile("[^a-zA-Z0-9.-]+")
	if err != nil {
		panic(err)
	}

	var groupID = reg.ReplaceAllString(parts[0], "")
	var artifactID = reg.ReplaceAllString(parts[1], "")
	var version = reg.ReplaceAllString(parts[2], "")

	var badG = strings.ContainsAny(groupID, " ")
	var badA = strings.ContainsAny(artifactID, " ")
	var badV = strings.ContainsAny(version, " ")

	if badG || badA || badV {
		return gav, errors.New("Could not parse coordinate. They had spaces: " + pText)
	}

	gav.GroupID = groupID
	gav.ArtifactID = artifactID
	gav.Version = version
	gav.Extension = "jar"
	gav.Classifier = ""
	

	return gav, nil
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