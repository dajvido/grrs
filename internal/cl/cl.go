package cl

import (
	"path/filepath"

	"github.com/urfave/cli/v2"
)

type Args struct {
	Pattern string
	Path    string
}

func InitArgs(cliArgs cli.Args) (*Args, error) {
	args := Args{Pattern: cliArgs.Get(0)}
	path, err := filepath.Abs(cliArgs.Get(1))
	args.Path = path
	return &args, err
}
