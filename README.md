# gfpush

`gfpush` is an uncomplicated command-line interface (CLI) utility designed to expedite the process of executing `git commit && git push origin <branch_name>`. It additionally provides support for the [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) specification.

[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-%23FE5196?logo=conventionalcommits&logoColor=white)](https://conventionalcommits.org)

## Installation

- Clone repository.
- Build binary with `make build`. Make sure `gcc` and `shc` is installed.
- Relocate `gfpush` to a directory listed in your $PATH. Example: `mv gfpush /usr/local/bin/`.

## Usage

- Include the required file(s) in the staging area using `git add`, or use `gfpush -a`, which functions similarly to `git commit -a`.
- Specify the commit message type using the `-t` flag and provide the commit message using `-m`.
- Include the scope of the commit message by specifying it with the `-s` flag. (Optional).
- To denote a *BREAKING CHANGE*, utilize the `-e` flag, which appends an exclamation mark (!) to the commit message. (Optional).

## Examples

| Command                                      | Commit message                       |
| :------------------------------------------- | :----------------------------------- |
| `gfpush -b -m 'Add OAuth2 via Keycloak'`   | branch_name: Add OAuth2 via Keycloak |
| `gfpush -t 5 -m 'Add OAuth2 via Keycloak'` | feat: Add OAuth2 via Keycloak        |
| `gfpush -e -t 4 -m 'Add Examples unit'`    | docs!: Add Examples unit             |
| `gfpush -e -t 5 -s api -m 'Add metrics'`   | feat(api)!: Add metrics              |

### Accepted `-t` flag values

| Value |   Type   | Description                                                                                      |
| :---: | :------: | :----------------------------------------------------------------------------------------------- |
|   1   |  build  | Changes that affect the build system or external dependencies                                    |
|   2   |  chore  | Changes that do not relate to a fix or feature and don't modify src or test files                |
|   3   |    ci    | Continuous integration related                                                                   |
|   4   |   docs   | Updates to documentation                                                                         |
|   5   |   feat   | New feature is introduced with the changes                                                       |
|   6   |   fix   | Bug fix has occurred                                                                             |
|   7   |   perf   | Performance improvements                                                                         |
|   8   | refactor | Refactored code that neither fixes a bug nor adds a feature                                      |
|   9   |  revert  | Reverts a previous commit                                                                        |
|  10  |  style  | Changes that do not affect the meaning of the code (white-space, missing semi-colons, and so on) |
|  11  |   test   | Including new or correcting previous tests                                                       |
