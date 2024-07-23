package runner_test

import (
	"bytes"
	"context"
	"os"
	"testing"

	"github.com/ooaklee/actions/go-example/go/internal/config"
	"github.com/ooaklee/actions/go-example/go/internal/runner"
	githubactions "github.com/sethvargo/go-githubactions"
	"github.com/stretchr/testify/assert"
)

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
				os.Setenv("GITHUB_WORKSPACE", "test/dir")
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
				os.Setenv("GITHUB_WORKSPACE", "")
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
			}

			err = runner.InvokeAction(ctx, cfg)
			if err != nil {
				assert.Equal(t, test.expectedError, err)
			}

			assert.Equal(t, test.expectedOutput, actionLog.String())
		})
	}
}
