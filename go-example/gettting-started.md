# Getting Started

This is a template for a Go-base GitHub Action project. It uses a combination of patterns taken from https://full-stack.blend.com/how-we-write-github-actions-in-go.html mixed in with some of the flows I've come across.

> Note, unless specified it's assumed that the any reference codeblocks will be ran from the Go actions' source code's root directory (./src)

## Usage

To test this action locally, you can run the following command:

```sh
env \
  'GITHUB_API_URL=https://api.github.com' \
  'GITHUB_REPOSITORY=blend/repo-that-uses-an-action' \
  "GITHUB_WORKSPACE=$(pwd)" \
  'INPUT_NAME=john' \
  'INPUT_REPETITION=5' \
  go run main.go
```

### Formatting code

You can use the following command to format the code:

```sh
go fmt  ./...
```

### Testing code

You can use the following command to format the code:

```sh
go test -v  ./...
```

### Building binary

There are two methods of building the binary. It can be done with native go build command.


```sh
# For Linux AMD64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags=\"-w -s\" -o dist/action-amd64 main.go

# For Linux ARM64
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -installsuffix cgo -ldflags=\"-w -s\" -o dist/action-arm64 main.go
```

**or**

It can be done using the package.json script (this is done automatically on push)

```sh
npm run package
```
