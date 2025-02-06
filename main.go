package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const version = "1.6.0"

var (
	message     string
	scope       string
	commitType  int
	flagBranch  bool
	flagAll     bool
	flagExclaim bool
	showHelp    bool
	showVersion bool

	commitTypes = []string{
		"", "build", "chore", "ci", "docs",
		"feat", "fix", "perf", "refactor",
		"revert", "style", "test",
	}
)

func main() {
	// Define CLI flags
	flag.BoolVar(&flagAll, "a", false, "Automatically stage modified and deleted files.")
	flag.BoolVar(&flagBranch, "b", false, "Include branch name in commit message.")
	flag.BoolVar(&flagExclaim, "e", false, "Add (!) to the commit convention.")
	flag.StringVar(&message, "m", "", "Commit message.")
	flag.StringVar(&scope, "s", "", "Scope of commit.")
	flag.IntVar(&commitType, "t", 0, "Type of commit message.")
	flag.BoolVar(&showHelp, "h", false, "Show help.")
	flag.BoolVar(&showVersion, "v", false, "Show version.")
	flag.Parse()

	if showHelp {
		printHelp()
		os.Exit(0)
	}

	if showVersion {
		fmt.Printf("gfpush version %s\n", version)
		os.Exit(0)
	}

	if message == "" {
		log.Fatal("Error: Commit message cannot be empty.")
	}

	if flagBranch {
		branch := getCurrentBranch()
		executeGitCommit(branch + ": " + message)
	} else if commitType > 0 && commitType < len(commitTypes) {
		commitPrefix := fmt.Sprintf("%s%s%s: ", commitTypes[commitType], formatScope(scope), addExclamation(flagExclaim))
		executeGitCommit(commitPrefix + message)
	} else {
		fmt.Println("Invalid commit type. Please choose a value between 1 and 11.")
		printCommitTypes()
		os.Exit(1)
	}

	executeGitPush()
}

func printHelp() {
	fmt.Printf(`
NAME:
   gfpush - Git Fast Push

USAGE:
   gfpush [global options...]

VERSION:
   %s

GLOBAL OPTIONS:
   -a                  Automatically stage modified and deleted files.
   -b                  Include branch name in commit message.
   -e                  Add (!) to the commit convention.
   -m <msg>            Commit message (required).
   -s <value>          Scope of commit.
   -t <value>          Commit type (1-11):
                         1: build      2: chore    3: ci
                         4: docs       5: feat     6: fix
                         7: perf       8: refactor 9: revert
                         10: style     11: test
   -h                  Show help.
   -v                  Show version.
`, version)
}

func printCommitTypes() {
	fmt.Println(`
Commit Types:
  1: build      2: chore    3: ci
  4: docs       5: feat     6: fix
  7: perf       8: refactor 9: revert
  10: style     11: test
`)
}

func formatScope(scope string) string {
	if scope == "" {
		return ""
	}
	return fmt.Sprintf("(%s)", scope)
}

func addExclamation(exclaim bool) string {
	if exclaim {
		return "!"
	}
	return ""
}

func getCurrentBranch() string {
	out, err := exec.Command("git", "branch", "--show-current").Output()
	if err != nil {
		log.Fatalf("Error fetching branch name: %v", err)
	}
	return strings.TrimSpace(string(out))
}

func executeGitCommit(commitMessage string) {
	args := []string{"commit"}
	if flagAll {
		args = append(args, "-a")
	}
	args = append(args, "-m", commitMessage)

	fmt.Println("Committing file(s) status:")
	executeGitCommand("git", "status", "-s")
	fmt.Println()

	executeGitCommand("git", args...)
}

func executeGitPush() {
	fmt.Println("\nPushing changes...")
	executeGitCommand("git", "push", "-q", "origin")
	fmt.Println("Everything up-to-date 🚀")
}

func executeGitCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error running command '%s': %v", strings.Join(cmd.Args, " "), err)
	}
}
