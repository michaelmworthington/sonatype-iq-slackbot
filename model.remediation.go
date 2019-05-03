package main

import (
	"encoding/json"
)

//RemediationResponse to contain the response
type RemediationResponse struct {
	Remediation Remediation  `json:"remediation"`
}

//Remediation to contain the version recommendations
type Remediation struct {
	VersionChanges []VersionChange  `json:"versionChanges"`
	//PolicyWaivers  []PolicyWaiver `json:"policyWaivers"`
	//ComponentOverrides []ComponentOverride `json:"componentOverrides"`
}

//VersionChange to contain the data and type
type VersionChange struct {
	Data Data  `json:"data"`
	Type string `json:"type"`
}

//Data to contain the component
type Data struct {
	Component Component  `json:"component"`
}

//Component to contain the component identifier and hash
type Component struct {
	ComponentIdentifier ComponentIdentifier  `json:"componentIdentifier"`
	Hash string `json:"hash"`
}


//ToRemediationResponseJSON to be used for marshalling of RemediationResponse type
func (b RemediationResponse) ToRemediationResponseJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

//FromRemediationResponseJSON to be used for unmarshalling of RemediationResponse type
func FromRemediationResponseJSON(data []byte) RemediationResponse {
	book := RemediationResponse{}
	err := json.Unmarshal(data, &book)
    //err := mapstructure.Decode(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}