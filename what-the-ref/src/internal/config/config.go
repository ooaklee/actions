package config

import (
	"errors"
	"strings"

	githubactions "github.com/sethvargo/go-githubactions"
)

const (
	// DefaultActionsStorePathGitHubHostedRunner is the default path
	// for the action store for GitHub hosted runners
	DefaultActionsStorePathGitHubHostedRunner string = "/home/runner/work/_actions"

	// DefaultActionsStorePathNonGitHubHostedRunner is the default path
	// for the action store for non GitHub hosted runners i.e. runs-on
	DefaultActionsStorePathNonGitHubHostedRunner string = "/home/runner/work/_actions"

	// DefaultActionHomePath is the default path for the action's home path
	DefaultActionHomePath string = "/home/runner"
)

var (
	// ErrNoActionNameProvided is the error returned when no action name is given
	ErrNoActionNameProvided = errors.New("NoActionNameProvided")
)

type Config struct {

	// ActionName is the name of the action to find
	// in the workflow's action store
	ActionName string

	// ActionHomePathOverride is an optional override
	// that defines the path for the action's home path
	// will use the value for the environment variable
	// HOME if not set
	ActionHomePathOverride string

	// ActionFullActionsStorePathOverride is an optional override
	// that defines the path for the action's full actions store
	// (where referenced actions of the workflow are stored).
	// If set, ActionHomePathOverride will not be used even if set
	ActionFullActionsStorePathOverride string

	// Dynamically set values

	Action *githubactions.Action
}

func NewFromInputs(action *githubactions.Action) (*Config, error) {

	var (
		actionName                         string
		actionHomePathOverride             string
		actionFullActionsStorePathOverride string
	)

	// handle input for action name
	actionNameInput := trimStringIfNotEmpty(
		action.GetInput("action-name"),
	)
	if actionNameInput == "" {
		action.Fatalf("An action name must be provided")
	}
	actionName = actionNameInput

	// handle input for action full actions store path override
	actionFullActionsStorePathOverrideInput := trimStringIfNotEmpty(
		action.GetInput("action-full-actions-store-path-override"),
	)

	if actionFullActionsStorePathOverrideInput != "" {
		action.Infof("Using provided action full actions store path override: %s", actionFullActionsStorePathOverrideInput)
		actionFullActionsStorePathOverride = actionFullActionsStorePathOverrideInput

		// remove trailing slash if present
		if strings.HasSuffix(actionFullActionsStorePathOverride, "/") {
			action.Debugf("Removing trailing slash from action full actions store path override: %s", actionFullActionsStorePathOverride)
			actionFullActionsStorePathOverride = actionFullActionsStorePathOverride[:len(actionFullActionsStorePathOverride)-1]
		}
	}

	// handle input for action home path override
	actionHomePathOverrideInput := trimStringIfNotEmpty(
		action.GetInput("action-home-path-override"),
	)

	//  if both overrides are set, the action home path override will not be used
	if actionHomePathOverrideInput != "" && actionFullActionsStorePathOverride != "" {
		action.Warningf("The action home path override (%s) will not be used because a full actions store path override (%s) was provided and takes precedence.", actionHomePathOverrideInput, actionFullActionsStorePathOverrideInput)
		actionHomePathOverride = ""
	}

	if actionHomePathOverrideInput != "" && actionFullActionsStorePathOverride == "" {
		action.Infof("Using provided action home path override: %s", actionHomePathOverrideInput)
		actionHomePathOverride = actionHomePathOverrideInput

		// remove trailing slash if present
		if strings.HasSuffix(actionHomePathOverride, "/") {
			action.Debugf("Removing trailing slash from action home path override: %s", actionHomePathOverride)
			actionHomePathOverride = actionHomePathOverride[:len(actionHomePathOverride)-1]
		}
	}

	c := Config{
		ActionName:                         actionName,
		ActionHomePathOverride:             actionHomePathOverride,
		ActionFullActionsStorePathOverride: actionFullActionsStorePathOverride,
		Action:                             action,
	}

	return &c, nil
}

// trimStringIfNotEmpty trims any leading or trailing whitespace from the input string s.
// If s is an empty string, an empty string is returned.
func trimStringIfNotEmpty(s string) string {
	if s != "" {
		return strings.TrimSpace(s)
	}

	return ""
}
