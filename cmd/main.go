package main

import (
	"fmt"
	"github.com/neonima/fftw/pkg/file"
	"github.com/neonima/fftw/pkg/stringster"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "fftw",
		Usage:  "A dead simple tool to rename you file for smooth web access!",
		Action: Run,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "recursive",
				Value:       false,
				Aliases:     []string{"r"},
				DefaultText: "Use this flag is you want to rename files recursively",
			},
			&cli.StringSliceFlag{
				Name:        "extensions",
				Aliases:     []string{"e"},
				DefaultText: "Only rename the specified extensions",
				Usage:       "-e=.png,.pdf | rename only `.png` et `.pdf` file",
			},
			&cli.StringFlag{
				Name:        "source",
				Aliases:     []string{"s"},
				DefaultText: "source path where to rename files",
			},
			&cli.BoolFlag{
				Name:        "dry",
				Aliases:     []string{"d"},
				DefaultText: "shows all the file that will be modified with their new name",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func Run(c *cli.Context) error {
	recurse := c.Bool("recursive")
	extensions := c.StringSlice("extensions")
	src := c.String("source")
	dry := c.Bool("dry")
	if len(src) > 0 {
		if err := os.Chdir(src); err != nil {
			return err
		}
	}
	src = "."

	if dry {
		fmt.Println("Files about to change")
	}
	filesToChange, err := file.Walk(src, recurse, extensions...)
	if err != nil {
		return err
	}

	for _, fileToRename := range filesToChange {
		newName := stringster.RenameFile(fileToRename)
		fmt.Println(fileToRename, "=>", newName)
		if !dry {
			if err = os.Rename(fileToRename, newName); err != nil {
				return err
			}
		}
	}

	return nil
}