package config_test

import (
	"bytes"
	"testing"

	"github.com/ooaklee/actions/what-the-ref/internal/config"
	githubactions "github.com/sethvargo/go-githubactions"
	"github.com/stretchr/testify/assert"
)

func TestConfig_NewFromInputs(t *testing.T) {

	tests := []struct {
		name   string
		preRun func()
		envMap map[string]string

		expectedOutput string
		expectedConfig config.Config
		expectedError  error
	}{
		{
			name:   "successful - just action name config from input",
			preRun: func() {},
			envMap: map[string]string{
				"INPUT_ACTION-NAME":                             "ooaklee/actions/what-the-ref",
				"INPUT_ACTION-HOME-PATH-OVERRIDE":               "",
				"INPUT_ACTION-FULL-ACTIONS-STORE-PATH-OVERRIDE": "",
			},
			expectedConfig: config.Config{
				ActionName: "ooaklee/actions/what-the-ref",
			},
			expectedOutput: "",
			expectedError:  nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			actionLog := bytes.NewBuffer(nil)

			test.preRun()

			getenv := func(key string) string {
				return test.envMap[key]
			}

			action := githubactions.New(
				githubactions.WithWriter(actionLog),
				githubactions.WithGetenv(getenv),
			)

			cfg, inputsErr := config.NewFromInputs(action)
			if inputsErr != nil {
				assert.Equal(t, test.expectedError, inputsErr)
			}

			if inputsErr == nil {
				assert.NotNil(t, cfg.Action)
				assert.Equal(t, test.expectedOutput, actionLog.String())

				// Make config's action nil for comparison
				cfg.Action = nil

				assert.Equal(t, test.expectedConfig, *cfg)
			}

		})
	}
}
