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

// InvokeAction runs the action behaviour using the parsed config and cancellation context.
func InvokeAction(ctx context.Context, cfg *config.Config) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	githubActionWorkingDir := os.Getenv("GITHUB_WORKSPACE")

	if githubActionWorkingDir == "" {
		cfg.Action.Errorf("GITHUB_WORKSPACE not found")
		return ErrGitHubWorkspaceEnvVarIsMissing
	}

	for i := 0; i < cfg.Repetition; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		cfg.Action.Infof("Hello, %s", cfg.Name)
	}

	cfg.Action.Infof("\n...I hope you are having a great day!")

	return nil

}
