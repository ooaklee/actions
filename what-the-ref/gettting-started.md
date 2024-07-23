# Getting Started

This is a GitHub Action that allows users to find the ref (sha/branch & path relative to workflow) of a named action that is utilised in a workflow. By using this action, users will be able to get the git reference, path and more.


> Note, unless specified it's assumed that the any reference codeblocks will be ran from the Go actions' source code's root directory (./src)

## Usage

To test this action locally, you can run the following command:

```sh
env \
  'GITHUB_API_URL=https://api.github.com' \
  "GITHUB_WORKSPACE=$(pwd)" \
  'GITHUB_REPOSITORY=ooaklee/actions/what-the-ref' \
  'INPUT_ACTION-NAME=ooaklee/actions/what-the-ref' \
  'INPUT_ACTION-HOME-PATH-OVERRIDE=' \
  'INPUT_ACTION-FULL-ACTIONS-STORE-PATH-OVERRIDE=' \
  go run main.go
```

> Note, the `HOME` var causes a clash will review this later

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
