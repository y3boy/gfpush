package main

import (
	"os"
	"log"
	"errors"
	"github.com/urfave/cli/v2"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

const version = "0.0.1"

func printCommitType() string {
	return `1: build - changes that affect the build system or external dependencies
	2: chore - changes that do not relate to a fix or feature and don't modify src or test files 
	3: ci - continuous integration related
	4: docs - updates to documentation 
	5: feat - a new feature is introduced with the changes
	6: fix - a bug fix has occurred
	7: perf - performance improvements
	8: refactor - refactored code that neither fixes a bug nor adds a feature
	9: revert - reverts a previous commit
	10: style - changes that do not affect the meaning of the code (white-space, missing semi-colons, and so on)
	11: test - including new or correcting previous tests`
}

func main() {
	commitType := map[string]string{
		"1": "build",
		"2": "chore",
		"3": "ci",
		"4": "docs",
		"5": "feat",
		"6": "fix",
		"7": "perf",
		"8": "refactor",
		"9": "revert",
		"10": "style",
		"11": "test",
	}

	// setup version flag
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "print-version",
		Aliases: []string{"v"},
		Usage:   "print gfpush version",
	}

	app := &cli.App{
		Name:                 "gfpush",
		Usage:                "commit and push faster",
		EnableBashCompletion: true,
		Version:              version,
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Anushervon Nabiev",
				Email: "nabievanush1@gmail.com",
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "Tell the command to automatically stage files that have been modified and deleted,\n\tbut new files you have not told Git about are not affected.",
			},
			&cli.BoolFlag{
				Name:    "exclamation-mark",
				Aliases: []string{"e"},
				Usage:   "Add '!' to convention.",
			},
			&cli.StringFlag{
				Name:    "message",
				Aliases: []string{"m"},
				Usage:   "Use the given `<msg>` as the commit message.",
			},
			&cli.StringFlag{
				Name:    "scope",
				Aliases: []string{"s"},
				Usage:   ("Scope of commit."),
			},
			&cli.StringFlag{
				Name:    "type",
				Aliases: []string{"t"},
				Usage:   ("Type of commit message.\n\t" + printCommitType()),
			},
		},
		Action: func(ctx *cli.Context) error {
			commitMessage := ""
			
			if _, ok := commitType[ctx.String("type")]; ok {
				commitMessage += commitType[ctx.String("type")]
			} else {
				return errors.New("gfpush: type of commit message not found")
			}

			if _, ok := commitType[ctx.String("scope")]; ok {
				commitMessage += "(" + ctx.String("scope") + ")"
			}
			
			if _, ok := commitType[ctx.String("exclamation-mark")]; ok {
				commitMessage = "!"
			}
			
			if len(ctx.String("message")) > 0 {
				commitMessage += ": " + ctx.String("message")
			} else {
				return errors.New("gfpush: commit message not found")
			}

			curr_path, err := os.Getwd()
			if err != nil {
				log.Println(err)
			}

			r, err := git.PlainOpen(curr_path)
			if err != nil {
				log.Println(err)
			}

			w, err := r.Worktree()
			if err != nil {
				log.Println(err)
			}

			w.Commit(commitMessage, &git.CommitOptions{
				Author: &object.Signature{
					Name:  "",
				},
			})
			r.Push(&git.PushOptions{})

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
