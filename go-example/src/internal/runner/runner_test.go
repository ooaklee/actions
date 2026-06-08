package runner_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/ooaklee/actions/go-example/go/internal/config"
	"github.com/ooaklee/actions/go-example/go/internal/runner"
	githubactions "github.com/sethvargo/go-githubactions"
	"github.com/stretchr/testify/assert"
)

// TestRunner_InvokeAction verifies workspace validation and greeting output behaviour.
func TestRunner_InvokeAction(t *testing.T) {

	tests := []struct {
		name   string
		preRun func()
		envMap map[string]string

		expectedOutput string
		expectedError  error
	}{
		{
			name: "successfully invoke action",
			preRun: func() {
			},
			envMap: map[string]string{
				"INPUT_NAME":       "john",
				"INPUT_REPETITION": "4",
			},
			expectedOutput: `Hello, john
Hello, john
Hello, john
Hello, john

...I hope you are having a great day!
`,
			expectedError: nil,
		},
		{
			name: "failed no workspace env detected",
			preRun: func() {
			},
			envMap: map[string]string{
				"INPUT_NAME":       "john",
				"INPUT_REPETITION": "4",
			},
			expectedOutput: "::error::GITHUB_WORKSPACE not found\n",
			expectedError:  runner.ErrGitHubWorkspaceEnvVarIsMissing,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			ctx := context.Background()
			actionLog := bytes.NewBuffer(nil)

			if test.expectedError == runner.ErrGitHubWorkspaceEnvVarIsMissing {
				t.Setenv("GITHUB_WORKSPACE", "")
			} else {
				t.Setenv("GITHUB_WORKSPACE", "test/dir")
			}

			test.preRun()

			getenv := func(key string) string {
				return test.envMap[key]
			}

			action := githubactions.New(
				githubactions.WithWriter(actionLog),
				githubactions.WithGetenv(getenv),
			)

			cfg, err := config.NewFromInputs(action)
			if err != nil {
				assert.Equal(t, test.expectedError, err)
				return
			}

			err = runner.InvokeAction(ctx, cfg)
			assert.Equal(t, test.expectedError, err)

			assert.Equal(t, test.expectedOutput, actionLog.String())
		})
	}
}
