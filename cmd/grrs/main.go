package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dajvido/grrs/internal/cl"
	"github.com/dajvido/grrs/internal/matches"
	"github.com/urfave/cli/v2"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

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
	check(err)
}
