package cli

import (
	"github.com/DataDrake/cli-ng/v2/cmd"
	"github.com/buildsi/codegen/generate"
	"github.com/buildsi/codegen/utils"
	"path/filepath"
)

// Args and flags for generate
type GenArgs struct {
	ConfigFile []string `zero:"true" desc:"A codegen.yaml to parse"`
}
type GenFlags struct{}

// Parser looks at symbols and ABI in Go
var Generator = cmd.Sub{
	Name:  "gen",
	Alias: "g",
	Short: "generate code from a codegen.yaml",
	Flags: &GenFlags{},
	Args:  &GenArgs{},
	Run:   RunGen,
}

func init() {
	cmd.Register(&Generator)
}

// RunParser reads a file and creates a corpus
func RunGen(r *cmd.Root, c *cmd.Sub) {
	args := c.Args.(*GenArgs)

	// If no config provided, assume in the PWD
	if len(args.ConfigFile) == 0 {
		args.ConfigFile = []string{filepath.Join(utils.GetPwd(), "codegen.yaml")}
	}
	generate.Generate(args.ConfigFile[0])
}
