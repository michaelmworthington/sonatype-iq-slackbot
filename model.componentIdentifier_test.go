package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

//   -d '{
//     "componentIdentifier": {
//         "format": "maven",
//         "coordinates": {
//             "artifactId": "log4j-core",
//             "classifier": "",
//             "extension" : "jar",
//             "groupId": "org.apache.logging.log4j",
//             "version": "2.5"
//         }
//     }
// }'
func TestToJSON(t *testing.T) {
	book := CI{ComponentIdentifier: 
				ComponentIdentifier{ Format: "maven",
									Coordinates: MavenCoordinate{ GroupID: "org.apache.logging.log4j",
																	ArtifactID: "log4j-core",
																	Classifier: "",
																	Extension : "jar",
																	Version: "2.5",
																},
									},
			}

	json := book.ToComponentIdentifierJSON()

	assert.Equal(t, 
				`{"componentIdentifier":{"format":"maven","coordinates":{"artifactId":"log4j-core","classifier":"","extension":"jar","groupId":"org.apache.logging.log4j","version":"2.5"}}}`, 
				string(json), 
				"JSON marshalling wrong.")
}

func TestFromJSON(t *testing.T) {
	json := []byte(`{"componentIdentifier": {"format": "maven","coordinates": {"artifactId": "log4j-core","classifier": "","extension" : "jar","groupId": "org.apache.logging.log4j","version": "2.5"}}}`)
	book := FromComponentIdentifierJSON(json)

	assert.Equal(t, 
				CI{ComponentIdentifier: 
					ComponentIdentifier{ Format: "maven",
										Coordinates: MavenCoordinate{ GroupID: "org.apache.logging.log4j",
																		ArtifactID: "log4j-core",
																		Classifier: "",
																		Extension : "jar",
																		Version: "2.5",
																	},
										},
				}, 
				book, 
				"JSON unmarshalling wrong.")
}

// func TestLookupAllVersions(t *testing.T) {
// 	lookupAllVersions("log4j-core")
// }