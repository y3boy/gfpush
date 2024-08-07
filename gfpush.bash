#!/bin/bash

usage() {
   echo "Try 'gfpush -h' for more information"
}

VERSION="1.2.0"

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
   -a                  Tell the command to automatically stage files that have been modified and deleted,
                       but new files you have not told Git about are not affected.
   -e                  Add (!) to convention.
   -m <msg>            Use the given <msg> as the commit message.
   -s value            Scope of commit.
   -t value            Type of commit message:
                        1: build - changes that affect the build system or external dependencies
                        2: chore - changes that do not relate to a fix or feature and dont modify src or test files 
                        3: ci - continuous integration related
                        4: docs - updates to documentation 
                        5: feat - a new feature is introduced with the changes
                        6: fix - a bug fix has occurred
                        7: perf - performance improvements
                        8: refactor - refactored code that neither fixes a bug nor adds a feature
                        9: revert - reverts a previous commit
                        10: style - changes that do not affect the meaning of the code (white-space, missing semi-colons, and so on)
                        11: test - including new or correcting previous tests
   -h                  Show help
   -v                  Print gfpush version.
   
EXAMPLE:
    gfpush -t 5 -m 'Add OAuth2 via Keycloak' --> feat: Add OAuth2 via Keycloak
    gfpush -e -t 4 -m 'Add Examples unit'    --> doc!: Add Examples unit 
    gfpush -e -t 5 -s api -m 'Add metrics'   --> feat(api)!: Add metrics
"

declare TYPE=( [1]="build" [2]="chore" [3]="ci" [4]="docs" 
               [5]="feat" [6]="fix" [7]="perf" [8]="refactor"
               [9]="revert" [10]="style" [11]="test" )


while getopts "ahm:st:v" options; do
   case "$options" in
   a)
      # git commit <-a> flag
      FLAG_ALL="-a"
   ;;
   e)
      # exclamation mark
      FLAG_E_MARK="!" 
   ;;
   h)
      echo "$HELP"
      exit 0
   ;;
   m)
      MESSAGE="$OPTARG"
   ;;
   s) 
      SCOPE="$OPTARG"
   ;;
   t)
      COMMIT_TYPE="$OPTARG"
   ;;
   v)
      echo "gfpush version $VERSION"
      exit 0
   ;;
   *)
      usage
      exit 1
   ;;
   esac
done

BRANCH=$(git branch 2> /dev/null | sed -e '/^[^*]/d' -e 's/* \(.*\)/\1/')

if [[ "$COMMIT_TYPE" -le "0" ]]  || [[ "$COMMIT_TYPE" -ge "12" ]] || [[ ! -n "$COMMIT_TYPE" ]]; then
   echo "Make sure to provide correct commit message type."
   echo -e "1: build   2: chore     3: ci\n4: docs    5: feat      6: fix\n7: perf    8: refactor  9: revert\n10: style  11: test"
   usage
   exit 1
fi

if [ -z "$MESSAGE" ]; then
   echo "Make sure commit message is not empty."
   usage
   exit 1
fi

echo "Commiting file(s) status:";
git status -s;
echo;

eval "git commit $FLAG_ALL -m '${TYPE[$COMMIT_TYPE]}$SCOPE$FLAG_E_MARK: $MESSAGE'";

if [ "$?" -gt 0 ]; then
   exit 1
fi

echo;

eval "git push -q origin $BRANCH";

if [[ "$?" -gt 0 ]]; then
   exit 1
fi

echo "Everything up-to-date 🚀";

exit 0
