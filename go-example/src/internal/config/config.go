package config

import (
	"errors"
	"strconv"

	githubactions "github.com/sethvargo/go-githubactions"
)

var (
	// ErrNameInputIsMissing is returned when the required name input is empty.
	ErrNameInputIsMissing = errors.New("NameInputIsMissing")

	// ErrInvalidRepetitionInput is returned when repetition cannot be parsed as an integer.
	ErrInvalidRepetitionInput = errors.New("InvalidRepetitionInput")

	// ErrRepetitionMustBePositive is returned when repetition is less than one.
	ErrRepetitionMustBePositive = errors.New("RepetitionMustBePositive")
)

// DefaultRepetition is used when the optional repetition input is not provided.
const DefaultRepetition = 1

// Config contains the parsed GitHub Action inputs and the action logger/client.
type Config struct {
	Name       string
	Repetition int
	Action     *githubactions.Action
}

// NewFromInputs reads GitHub Action inputs, validates them, and returns a Config.
func NewFromInputs(action *githubactions.Action) (*Config, error) {
	name := action.GetInput("name")
	if name == "" {
		action.Errorf("The name input was not provided")
		return nil, ErrNameInputIsMissing
	}

	// The action metadata provides a default, but keeping a code-level default
	// makes local runs and tests behave the same way.
	repetitionInput := action.GetInput("repetition")
	if repetitionInput == "" {
		action.Debugf("The repetition input was not provided, using default value of %d", DefaultRepetition)
		repetitionInput = strconv.Itoa(DefaultRepetition)
	}

	repetition, err := strconv.Atoi(repetitionInput)
	if err != nil {
		action.Errorf("Cannot convert the 'repetition' input (%s) to an int", repetitionInput)
		return nil, ErrInvalidRepetitionInput
	}

	if repetition < 1 {
		action.Errorf("The 'repetition' input must be greater than 0")
		return nil, ErrRepetitionMustBePositive
	}

	c := Config{
		Name:       name,
		Repetition: repetition,
		Action:     action,
	}
	return &c, nil
}
