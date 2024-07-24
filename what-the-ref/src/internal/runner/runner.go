package runner

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/ooaklee/actions/what-the-ref/internal/config"
)

const (
	// DefaultActionsStorePathGitHubHostedRunnerTmpl is the template of the
	// default path for the action store for GitHub hosted runners
	DefaultActionsStorePathGitHubHostedRunnerTmpl string = "%s/work/_actions"

	// DefaultActionsStorePathNonGitHubHostedRunnerTmpl is the tempate of the
	// default path for the action store for non GitHub hosted runners i.e. runs-on
	DefaultActionsStorePathNonGitHubHostedRunnerTmpl string = "%s/work/_actions"

	// DefaultActionHomePath is the default path for the action's home path
	DefaultActionHomePath string = "/home/runner"
)

var (
	// ErrGitHubWorkspaceEnvVarIsMissing is returned when the GitHub Workspace
	// environment variable is not set.
	ErrGitHubWorkspaceEnvVarIsMissing = errors.New("GitHubWorkspaceEnvVarIsMissing")

	// ErrActionStorePathNotFound is returned when the action store
	// path is not found.
	ErrActionStorePathNotFound = errors.New("ActionStorePathNotFound")

	// ErrActionHomePathNotFound is returned when the action home
	// path is not found.
	ErrActionHomePathNotFound = errors.New("ActionHomePathNotFound")

	// ErrValidTargetActionRefNotFound is returned when the action store
	// does not have a .completed version of the target action
	ErrValidTargetActionRefNotFound = errors.New("ValidTargetActionRefNotFound")
)

func InvokeAction(ctx context.Context, cfg *config.Config) error {
	var (
		githubActionWorkingDir string = os.Getenv("GITHUB_WORKSPACE")

		//
		isStorePathValid   bool = false
		possibleStorePaths []string
		actionStorePath    string
		actionHomePath     string

		finalFullActionsStorePath string
		finalActionRef            string
	)

	if githubActionWorkingDir == "" {
		cfg.Action.Errorf("GITHUB_WORKSPACE not found")
		return ErrGitHubWorkspaceEnvVarIsMissing
	}

	//// Get possible store paths
	// check if the action full actions store path override is set, if so, set as action store path
	if cfg.ActionFullActionsStorePathOverride != "" {
		possibleStorePaths = []string{cfg.ActionFullActionsStorePathOverride}
	}

	if cfg.ActionFullActionsStorePathOverride == "" {

		if cfg.ActionHomePathOverride != "" {
			cfg.Action.Infof("Using provided action home path override: %s", cfg.ActionHomePathOverride)
			actionHomePath = cfg.ActionHomePathOverride
		}

		if cfg.ActionHomePathOverride == "" {

			// attempt to find the action home path from the environment variable HOME
			actionHomePath := os.Getenv("HOME")

			if actionHomePath == "" {
				cfg.Action.Errorf("the HOME path cannot be found and it must be set, or a full actions store path override must be provided")
				cfg.Action.Debugf("the HOME path can be set by:\n  - setting the env HOME\n  - setting the input 'action-home-path-override'")
				return ErrActionHomePathNotFound
			}
		}

		possibleStorePaths = []string{
			// using action home path, create the default path for a GitHub hosted runner's action store
			fmt.Sprintf(
				DefaultActionsStorePathGitHubHostedRunnerTmpl,
				actionHomePath,
			),

			// using action home path, create the default path for a self-host hosted runner's action store
			fmt.Sprintf(
				DefaultActionsStorePathNonGitHubHostedRunnerTmpl,
				actionHomePath,
			),
		}
	}

	//// Validate action store path (ctx, possibleStorePaths) ( storePath string, error)
	// validate at least one of the possible store paths exist and is valid
	for _, possibleStorePath := range possibleStorePaths {
		cfg.Action.Infof("Checking if action store path exists: %s", possibleStorePath)

		if _, err := os.Stat(possibleStorePath); err == nil {
			isStorePathValid = true
			actionStorePath = possibleStorePath
			break
		}
	}

	if !isStorePathValid {
		cfg.Action.Errorf("the full action store path is not valid!")
		cfg.Action.Warningf("the full action store path can be set by:\n  - ensuring the correct HOME is set for your runner set up\n  - setting the input 'action-full-actions-store-path-override' to the correct path.\nBy default, each runner has the deafult path for the action store:\n - Github hosted runner: %s\n - Github self-hosted runner: %s",
			fmt.Sprintf(
				DefaultActionsStorePathGitHubHostedRunnerTmpl,
				DefaultActionHomePath,
			),
			fmt.Sprintf(
				DefaultActionsStorePathNonGitHubHostedRunnerTmpl,
				DefaultActionHomePath,
			),
		)
		return ErrActionStorePathNotFound
	}

	//// ValidateTargetActionExistsInActionStore(ctx, actionStorePath, actionName) ( actionRef string, actionFullPathInActionStore string, error)
	// validate that the action store path contains the action name and it is "ready" (contains ".completed" file)

	// check all files in the action store path and see if any have the suffix ".completed"
	actionStorePathForTargetActionContents, err := os.ReadDir(actionStorePath + "/" + cfg.ActionName)
	if err != nil {
		cfg.Action.Errorf("Error reading action store path: %s", err)
		return err
	}

	// holder for all the versions of the target action stored in the action store
	// will need to append to actionStorePath to get the full path to the stored action
	validTargetActionStoredVersions := []string{}
	detectedContentInActionStore := []string{}

	// check the contents of the action store's directory for target action
	for _, actionStorePathContent := range actionStorePathForTargetActionContents {

		if !actionStorePathContent.IsDir() && strings.HasSuffix(actionStorePathContent.Name(), ".completed") {
			validTargetActionStoredVersions = append(validTargetActionStoredVersions, strings.ReplaceAll(actionStorePathContent.Name(), ".completed", ""))
			continue
		}

		detectedContentInActionStore = append(detectedContentInActionStore, actionStorePath+"/"+actionStorePathContent.Name())
	}

	if len(validTargetActionStoredVersions) == 0 {
		cfg.Action.Errorf("No valid target action found in action store path: %s", actionStorePath)
		cfg.Action.Debugf("%d other detected content in action store path of the target action included:%s", len(detectedContentInActionStore), strings.Join(detectedContentInActionStore, "\n  - "))
		return ErrValidTargetActionRefNotFound
	}

	// todo: improve later to have better selection process
	// for now, just use the first valid target action stored version
	selectedTargetActionStoredVersion := validTargetActionStoredVersions[0]
	actionFullPathInActionStore := fmt.Sprintf("%s/%s/%s", actionStorePath, cfg.ActionName, selectedTargetActionStoredVersion)

	finalActionRef = selectedTargetActionStoredVersion
	finalFullActionsStorePath = actionFullPathInActionStore

	// ORIGINAL PLAN OF ACTION
	// otherwise check if the action home path override is set, if not, check for the HOME environment variable
	// if neither is set, error out and tell use that the action home path is not set and needs to be set
	// if value can be found for action home path, append into both  store path templates and see which path exists

	// if neither store path exists, error out and tell user that the action full store path is not set and cannot be found
	// if actionStorePath == "" && !isStorePathValid {
	// 	cfg.Action.Errorf("No action store path found")
	// 	return ErrActionStorePathNotFound
	// }

	// using the path that  exists, look for a file that has the suffix .completed
	// if the file cannot be found, error out and tell user that no valid action has not been completed (output dir in debug) and exit with error

	// if file it found, task it's name and assign as the ref output
	// verify the ref can be accessed and then use that full path and assign as the path output (making sure trailing slash is removed)

	// set the outputs
	cfg.Action.SetOutput("ref", finalActionRef)
	cfg.Action.SetOutput("path", finalFullActionsStorePath)

	// Summary exist of what was found
	cfg.Action.Infof("\"What the ref\" has determined that the candidate details for the specified action (%s) are as follows:\n  - Ref: %s\n  - Store Path: %s",
		cfg.ActionName,
		finalActionRef,
		finalFullActionsStorePath,
	)

	return nil

}
