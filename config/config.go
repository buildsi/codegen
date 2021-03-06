package config

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/buildsi/codegen/utils"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

// A codegen.yaml config file, with types defined for better control
type Conf struct {
	Rendering `yaml:"generate,omitempty"`
	Root      string `yaml:"root,omitempty"`
}

// A rendering holds parameters for a rendering session
type Rendering struct {
	Language string            `yaml:"language,omitempty"`
	Files    []string          `yaml:"files,omitempty"`
	Type     string            `yaml:"type,omitempty"`
	Renders  map[string]Render `yaml:"render,omitempty"`
}

// A render is a specific type and parameter settings for a thing like a function
type Render struct {
	Type       string            `yaml:"type,omitempty"`
	Numeric    bool              `yaml:"numeric,omitempty"`
	Parameters ParameterSettings `yaml:"parameters,omitempty"`
}

// A Parameter defines how the parameters should be generated
type ParameterSettings struct {
	Min      int      `yaml:"min,omitempty"`
	Max      int      `yaml:"max,omitempty"`
	Exact    int      `yaml:"exact,omitempty"`
	Pointers bool     `yaml:"pointers"`
	Types    []string `yaml:"types,omitempty"`
}

// read the config and return a config type
func readConfig(yamlStr []byte) Conf {

	// First unmarshall into generic structure
	var data map[string]interface{}
	err := yaml.Unmarshal(yamlStr, &data)
	if err != nil {
		log.Fatalf("Unmarshal: %v\n", err)
	}

	// A config can hold multiple keyed sections
	c := Conf{}

	// Load generation settings
	if item, ok := data["generate"]; ok {
		c.Rendering = loadRendering(item)
	}
	return c
}

// loadRendering loads the config section for a rendering preference (e.g., cpp)
func loadRendering(item interface{}) Rendering {
	rendering := Rendering{}
	mapstructure.Decode(item, &rendering)
	settings := item.(map[string]interface{})["render"]
	renders := map[string]Render{}
	mapstructure.Decode(settings, &renders)

	// Default needs to be true
	for key, render := range renders {

		// yes, this is really ugly!
		usePointer := settings.(map[string]interface{})[key].(map[string]interface{})["parameters"].(map[string]interface{})["pointers"]
		if usePointer == nil {
			render.Parameters.Pointers = true
		} else {
			render.Parameters.Pointers = usePointer.(bool)
		}
		renders[key] = render
	}
	rendering.Renders = renders
	return rendering
}

func Load(configFile string) Conf {

	// Ensure our config file exists!
	configFile = utils.AbsPath(configFile)
	if !utils.Exists(configFile) {
		log.Fatalf("%s does not exist.", configFile)
	}

	yamlContent, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	conf := readConfig(yamlContent)
	conf.Root = filepath.Dir(configFile)
	return conf
}
