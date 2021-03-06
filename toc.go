package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/yankeexe/toc-md/actions"
	"github.com/yankeexe/toc-md/utils"
	"log"
	"os"
	"path/filepath"
)

func main() {
	app := cli.NewApp()
	app.Name = "Table of contents Markdown."
	app.Usage = "Generate of inject table of contents to your markdown file"
	app.Commands = []*cli.Command{
		{
			Name:      "gen",
			Usage:     "Provide Markdown file name or path to the file.",
			UsageText: "toc gen <filename> | toc gen <filename> -i",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "inject",
					Usage:   "Inject TOC to file",
					Aliases: []string{"i"},
				},
			},

			Action: func(c *cli.Context) error {
				argCount := c.NArg()

				// Show help message based on arguments len.
				if argCount == 0 || argCount > 2 {
					return cli.ShowCommandHelp(c, "gen")
				}

				// Get argument value
				inject := c.Args().Get(1)
				fileLocation := c.Args().Get(0)
				extension := filepath.Ext(fileLocation)

				// File validations
				utils.ValidateFile(fileLocation, extension)

				file := utils.ReadFile(fileLocation)
				defer file.Close()

				// All the lines of file.
				var lines []string

				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					lines = append(lines, scanner.Text())
				}

				result := actions.GenerateToc(lines)

				// Inject to file or stdout.
				if inject != "" {
					actions.InjectToc(lines, result, file)
				} else {
					for _, item := range result {
						fmt.Println(item)
					}

				}
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
