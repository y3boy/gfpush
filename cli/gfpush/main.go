package main

const version = "0.0.1"

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	// commitType := map[int]string{
	// 	1: "fix", 
	// 	2: "feat",
	// 	3: "build",
	// 	4: "chore",
	// 	5: "ci",
	// 	6: "docs",
	// 	7: "style",
	// 	8: "refactor",
	// 	9: "perf",
	// 	10: "test",
	// }
	
	app := &cli.App{
		Name:  "gfpush",
		Usage: "commit and push faster",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "getversion",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("added task: ", cCtx.Args().First())
					return nil
				},
			},
	}

	app.Run(os.Args)
}