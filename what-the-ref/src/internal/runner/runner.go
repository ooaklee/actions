package runner

import (
	"context"
	"errors"
	"os"

	"github.com/ooaklee/actions/what-the-ref/internal/config"
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

	// Temp entrypoint output
	cfg.Action.Infof("What the ref has been invoked with the inputs: \n  - Action name:  %s\n  - Action home path override:  %s\n  - Action store path override:  %s",
		cfg.ActionName,
		cfg.ActionHomePathOverride,
		cfg.ActionFullActionsStorePathOverride,
	)

	return nil

}
