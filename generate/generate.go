package generate

import (
	"fmt"

	"github.com/buildsi/codegen/config"
	"github.com/buildsi/codegen/generate/cpp"
)

// Generate generates code based on a config file
func Generate(configFile string, outdir string, renderType string) {
	conf := config.Load(configFile)

	switch conf.Language {
	case "cpp":
		cpp.Generate(conf, outdir, renderType)
	default:
		fmt.Printf("%s is not a supported language", conf.Language)
	}
}
