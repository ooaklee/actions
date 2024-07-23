package main

import (
	"context"

	"github.com/ooaklee/actions/go-example/go/internal/config"
	"github.com/ooaklee/actions/go-example/go/internal/runner"
	githubactions "github.com/sethvargo/go-githubactions"
)

func run() error {
	ctx := context.Background()
	action := githubactions.New()

	cfg, err := config.NewFromInputs(action)
	if err != nil {
		return err
	}

	return runner.InvokeAction(ctx, cfg)
}

func main() {
	err := run()
	if err != nil {
		githubactions.Fatalf("%v", err)
	}
}
