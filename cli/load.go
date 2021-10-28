package cli

import (
	"github.com/DataDrake/cli-ng/v2/cmd"
	"github.com/buildsi/codegen/generate"
)

// Args and flags for load
type LoadArgs struct {
	JsonFile []string `zero:"true" desc:"A codegen.json to load"`
}
type LoadFlags struct{}

var Loader = cmd.Sub{
	Name:  "load",
	Alias: "load",
	Short: "load a codegen.json",
	Flags: &LoadFlags{},
	Args:  &LoadArgs{},
	Run:   RunLoad,
}

func init() {
	cmd.Register(&Loader)
}

func RunLoad(r *cmd.Root, c *cmd.Sub) {
	args := c.Args.(*LoadArgs)
	generate.Load(args.JsonFile[0])
}
