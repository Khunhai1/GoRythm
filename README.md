# GoRythm

## Usage

The `go run` command can be used to run the project.

```bash
$ go run ./cmd/main/main.go
```

## Web application

The project is hosted on [Github Pages](https://khunhai1.github.io/GoRythm/) using WebAssembly.

## Devcontainer

The devcontainer located in [.devcontainer](.devcontainer) creates a development environment for this project. It uses a ubuntu image with golang and all the pre-commit hooks installed.

The command `Devcontainer: Reopen in Container` in VSCode will open the project in the devcontainer.

It can't be used to run the project because of peripheral access issues.

The pre-commit hooks can be run with the following command.

```bash
$ pre-commit run --all-files
```

### Requirements

- Go 1.23.1
- VSCode with the devcontainer extension
- Docker
