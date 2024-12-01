#!/bin/bash

# Display usage instructions
usage() {
   echo "Try 'gfpush -h' for more information."
}

VERSION="1.5.0"

# Help message for the script
HELP="
NAME:
   gfpush - Git Fast Push

USAGE:
   gfpush [global options...]

VERSION:
   $VERSION

AUTHOR:
   Anushervon Nabiev <nabievanush1@gmail.com>

GLOBAL OPTIONS:
   -a                  Automatically stage modified and deleted files (new files are not affected).
   -b                  Add branch name to the commit message.
   -e                  Add (!) to the commit message convention.
   -m <msg>            Commit message.
   -s <scope>          Commit scope.
   -t <type>           Commit message type:
                          1: build    6: fix       11: test
                          2: chore    7: perf      
                          3: ci       8: refactor  
                          4: docs     9: revert    
                          5: feat    10: style     
   -h                  Show help.
   -v                  Print gfpush version.

EXAMPLES:
   gfpush -b -m 'Add OAuth2 via Keycloak'   --> branch_name: Add OAuth2 via Keycloak
   gfpush -t 5 -m 'Add OAuth2 via Keycloak' --> feat: Add OAuth2 via Keycloak
   gfpush -e -t 4 -m 'Add Examples unit'    --> docs!: Add Examples unit
   gfpush -e -t 5 -s api -m 'Add metrics'   --> feat(api)!: Add metrics
"

# Define commit types
declare -A TYPE=(
   [1]="build" [2]="chore" [3]="ci" [4]="docs"
   [5]="feat"  [6]="fix"   [7]="perf" [8]="refactor"
   [9]="revert" [10]="style" [11]="test"
)

# Initialize variables
FLAG_ALL=""
FLAG_BRANCH=false
FLAG_E_MARK=""
MESSAGE=""
SCOPE=""
COMMIT_TYPE=""
BRANCH=""

# Parse options
while getopts "abhm:s:t:v" options; do
   case "$options" in
      a) FLAG_ALL="-a" ;;                              # Enable auto-staging
      b) FLAG_BRANCH=true ;;                           # Use branch name in commit message
      e) FLAG_E_MARK="!" ;;                            # Add "!" to the commit message
      h) echo "$HELP"; exit 0 ;;                       # Display help
      m) MESSAGE="$OPTARG" ;;                          # Commit message
      s) SCOPE="($OPTARG)" ;;                          # Commit scope
      t) COMMIT_TYPE="$OPTARG" ;;                      # Commit type
      v) echo "gfpush version $VERSION"; exit 0 ;;     # Display version
      *) usage; exit 1 ;;                              # Invalid option
   esac
done

# Ensure a commit message is provided
if [ -z "$MESSAGE" ]; then
   echo "Error: Commit message cannot be empty."
   usage
   exit 1
fi

# Add branch name to the commit message if -b flag is used
if [[ "$FLAG_BRANCH" == true ]]; then
   BRANCH=$(git rev-parse --abbrev-ref HEAD 2>/dev/null)
   echo "Committing files with branch name:"
   git status -s
   echo
   git commit $FLAG_ALL -m "$BRANCH: $MESSAGE"

# Add commit type and scope to the message
elif [[ -n "$COMMIT_TYPE" && ${TYPE[$COMMIT_TYPE]+_} ]]; then
   echo "Committing files with type and scope:"
   git status -s
   echo
   git commit $FLAG_ALL -m "${TYPE[$COMMIT_TYPE]}$SCOPE$FLAG_E_MARK: $MESSAGE"

# Error handling for incorrect or missing options
else
   echo "Error: Provide a valid commit type or use the -b flag."
   echo -e "Commit types:\n1: build   2: chore    3: ci\n4: docs    5: feat     6: fix\n7: perf    8: refactor 9: revert\n10: style  11: test"
   usage
   exit 1
fi

# Push the changes
echo
git push -q origin $BRANCH
echo "Everything up-to-date 🚀"

exit 0
