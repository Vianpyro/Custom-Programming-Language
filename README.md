# Custom Programming Language

This template is used and expanded for use inside code practices for [DCI](https://github.com/dciets).

It's a custom interpreted language made with engineering practices such as test-driven development, and probably CI/CD in the future.

## Development Environment

This project includes a [Dev Container](https://containers.dev/) configuration for a consistent development environment. 
The dev container runs `golang:tip-alpine3.21`, which provides the latest Go development tools (*to this day*) in a lightweight Alpine Linux environment.

### Using the Dev Container

To use the dev container:

1. Install [Docker](https://www.docker.com/) and [VS Code](https://code.visualstudio.com/).
2. Install the [Dev Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) for VS Code.
3. Open the project in VS Code and, when prompted, reopen it inside the dev container.

This will set up the Go development environment automatically.

## Building the Project

To build the project, run:

```sh
go build -o custom-lang
```

This compiles the interpreter into an executable named `custom-lang`.

## Running the Project

To run the interpreter on a source file:

```sh
./custom-lang --main path/to/source-file
```
