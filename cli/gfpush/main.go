package main

import (
	"os"
	// "fmt"
	// "strconv"
	"github.com/urfave/cli/v2"
	// "github.com/go-git/go-git/v5"
)

const version = "0.0.1"

func printCommitType() string {
	return `1: build
	2: chore
	3: ci
	4: docs
	5: feat
	6: fix
	7: perf
	8: refactor
	9: revert
	10: style
	11: test`
}

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
	
	// setup version flag 
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "print-version",
		Aliases: []string{"v"},
		Usage:   "print gfpush version",
	}

	app := &cli.App{
		Name:  "gfpush",
		Usage: "commit and push faster",
		EnableBashCompletion: true,
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Anushervon Nabiev",
				Email: "nabievanush1@gmail.com",
			},
		},
		Version: version,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:     "all",
				Aliases:	[]string{"a"},
				Usage:       "Tell the command to automatically stage files that have been modified and deleted,\n\tbut new files you have not told Git about are not affected.",
			},
			&cli.StringFlag{
				Name:    "message",
				Aliases: []string{"m"},
				Usage:   "Use the given `<msg>` as the commit message.",
			},
			&cli.StringFlag{
				Name:    "type",
				Aliases: []string{"t"},
				Usage:   ("Type of commit message.\n\t" + printCommitType()),
			},
			&cli.StringFlag{
				Name:    "scope",
				Aliases: []string{"s"},
				Usage:   ("Scope of commit."),
			},
		},
	}
	app.Run(os.Args)
}