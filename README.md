# GopherTicTacToe

## Devcontainer

The devcontainer located in [.devcontainer](.devcontainer) creates a development environment for this project. It uses a ubuntu image with golang and all the pre-commit hooks installed.

### Usage

The command "Devcontainer: Reopen in Container" in VSCode will open the project in the devcontainer.

The go run command can be used to run the project.
```bash
$ go run .
```

The pre-commit hooks can be run with the following command.
```bash
$ pre-commit run --all-files
```

### Requirements

- VSCode with the devcontainer extension
- Docker
