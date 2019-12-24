package cl

import (
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

// Args represents arguments provided from command.
type Args struct {
	Pattern string // pattern for searching matches
	Path    string // path to the file to read
}

// InitArgs reads arguments from command and wrap them into a returned structure.
func InitArgs(cliArgs cli.Args) (*Args, error) {
	args := Args{Pattern: cliArgs.Get(0)}
	path, err := filepath.Abs(cliArgs.Get(1))
	if err != nil {
		return nil, errors.Wrap(err, path)
	}
	args.Path = path
	return &args, nil
}
