// codegen: generate code using Go

package main

import (
	"github.com/buildsi/codegen/cli"
)

// Run the codegen client against a config to generate code
func main() {
	cli.Root.Run()
}
