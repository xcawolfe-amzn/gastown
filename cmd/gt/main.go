// gt is the Gas Town CLI for managing multi-agent workspaces.
package main

import (
	"os"

	"github.com/xcawolfe-amzn/gastown/internal/cmd"
)

func main() {
	os.Exit(cmd.Execute())
}
