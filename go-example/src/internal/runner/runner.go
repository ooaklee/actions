package runner

import (
	"context"
	"errors"
	"os"

	"github.com/ooaklee/actions/go-example/go/internal/config"
)

var (
	// ErrGitHubWorkspaceEnvVarIsMissing is returned when the GitHub Workspace
	// environment variable is not set.
	ErrGitHubWorkspaceEnvVarIsMissing = errors.New("GitHubWorkspaceEnvVarIsMissing")
)

func InvokeAction(ctx context.Context, cfg *config.Config) error {
	var githubActionWorkingDir string = os.Getenv("GITHUB_WORKSPACE")

	if githubActionWorkingDir == "" {
		cfg.Action.Errorf("GITHUB_WORKSPACE not found")
		return ErrGitHubWorkspaceEnvVarIsMissing
	}

	for i := 0; i < cfg.Repetition; i++ {
		cfg.Action.Infof("Hello, %s", cfg.Name)
	}

	cfg.Action.Infof("\n...I hope you are having a great day!")

	return nil

}
