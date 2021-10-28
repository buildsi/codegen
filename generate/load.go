package generate

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// A generic load can be done across generators (output should be fairly consitent)
type FormalParam struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	IsSigned  bool   `json:"is_signed"`
	IsPointer bool   `json:"is_pointer"`
	Value     string `json:"value"`
}

// A Function holds a name and one or more formal params
type Function struct {
	Name         string        `json:"name"`
	FormalParams []FormalParam `json:"parameters,omitempty"`
}

// Load exported codegen.json into structure
func Load(jsonFile string) map[string]Function {

	content, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Printf("error reading %s #%v ", jsonFile, err)
	}

	funcs := map[string]Function{}
	err = json.Unmarshal(content, &funcs)
	if err != nil {
		log.Fatalf("Unmarshal: %v\n", err)
	}

	return funcs
}
