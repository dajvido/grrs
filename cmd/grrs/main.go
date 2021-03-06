package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dajvido/grrs/internal/cl"
	"github.com/dajvido/grrs/internal/matches"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

// check validate an error and log it.
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// main initialise cli and check for errors.
func main() {
	app := &cli.App{
		Name:  "grrs",
		Usage: "like a grep, but written in go",
		Action: func(c *cli.Context) error {
			args, err := cl.InitArgs(c.Args())
			check(err)
			lines, err := matches.InFile(args)
			check(err)
			for _, line := range lines {
				fmt.Println(line)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	check(errors.Wrap(err, strings.Join(os.Args, ", ")))
}
