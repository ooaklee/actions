package config_test

import (
	"bytes"
	"testing"

	"github.com/ooaklee/actions/go-example/go/internal/config"
	githubactions "github.com/sethvargo/go-githubactions"
	"github.com/stretchr/testify/assert"
)

// TestConfig_NewFromInputs verifies that action inputs are parsed, defaulted, and validated.
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
			name:   "successful - created base config from input",
			preRun: func() {},
			envMap: map[string]string{
				"INPUT_NAME":       "john",
				"INPUT_REPETITION": "4",
			},
			expectedConfig: config.Config{
				Name:       "john",
				Repetition: 4,
			},
			expectedOutput: "",
			expectedError:  nil,
		},
		{
			name:   "successful - created config with default repetition",
			preRun: func() {},
			envMap: map[string]string{
				"INPUT_NAME": "john",
			},
			expectedConfig: config.Config{
				Name:       "john",
				Repetition: config.DefaultRepetition,
			},
			expectedOutput: "::debug::The repetition input was not provided, using default value of 1\n",
			expectedError:  nil,
		},
		{
			name:   "failed - missing name input",
			preRun: func() {},
			envMap: map[string]string{
				"INPUT_REPETITION": "4",
			},
			expectedOutput: "::error::The name input was not provided\n",
			expectedError:  config.ErrNameInputIsMissing,
		},
		{
			name:   "failed - invalid repetition input",
			preRun: func() {},
			envMap: map[string]string{
				"INPUT_NAME":       "john",
				"INPUT_REPETITION": "many",
			},
			expectedOutput: "::error::Cannot convert the 'repetition' input (many) to an int\n",
			expectedError:  config.ErrInvalidRepetitionInput,
		},
		{
			name:   "failed - repetition must be positive",
			preRun: func() {},
			envMap: map[string]string{
				"INPUT_NAME":       "john",
				"INPUT_REPETITION": "0",
			},
			expectedOutput: "::error::The 'repetition' input must be greater than 0\n",
			expectedError:  config.ErrRepetitionMustBePositive,
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
				assert.Equal(t, test.expectedOutput, actionLog.String())
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
