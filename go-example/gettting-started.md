# Getting Started

This is a template for a Go-based GitHub Action project. It uses a combination of patterns taken from https://full-stack.blend.com/how-we-write-github-actions-in-go.html mixed in with some of the flows I've come across.

> Note, unless specified it's assumed that reference code blocks run from the Go action source code root directory (`./src`).

## Usage

To test this action locally, you can run the following command:

```sh
env \
  'GITHUB_API_URL=https://api.github.com' \
  'GITHUB_REPOSITORY=blend/repo-that-uses-an-action' \
  "GITHUB_WORKSPACE=$(pwd)/.." \
  'INPUT_NAME=john' \
  'INPUT_REPETITION=5' \
  go run main.go
```

## Adapting the template

When creating a new Go action from this example:

1. Update `action.yml` first so inputs and generated docs match the public API.
2. Add input parsing and validation in `src/internal/config`.
3. Put action behaviour in `src/internal/runner`.
4. Keep `invoke-binary.js` as the Node action entrypoint unless the binary layout changes.
5. Run `yarn test`, `yarn package`, and `yarn readme --update` from the repository root.

### Formatting code

You can use the following command to format the code:

```sh
go fmt  ./...
```

### Testing code

You can use the following command to test the code:

```sh
go test -v ./...
```

### Testing coverage

You can use the following command to generate a coverage report:

```sh
go test -coverprofile=cover.out ./...
go tool cover -html=cover.out
```

### Building binary

There are two methods of building the binary. It can be done with native go build command.

```sh
# For Linux AMD64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o ../dist/action-amd64 main.go

# For Linux ARM64
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -trimpath -ldflags="-s -w" -o ../dist/action-arm64 main.go
```

**or**

It can be done using the package.json script (this is done automatically on push)

```sh
yarn package
```
