package config

import (
	"strconv"

	"errors"

	githubactions "github.com/sethvargo/go-githubactions"
)

type Config struct {
	Name       string
	Repetition int
	Action     *githubactions.Action
}

func NewFromInputs(action *githubactions.Action) (*Config, error) {

	// handle input for how many times the greeting should
	// be repeated
	repetitionInput := action.GetInput("repetition")
	repetition, err := strconv.Atoi(repetitionInput)
	if err != nil {
		action.Fatalf("Can't convert the 'repetition' input (%s) to an int!", repetitionInput)
		return nil, errors.New("exit status 1")
	}

	c := Config{
		Name:       action.GetInput("name"),
		Repetition: repetition,
		Action:     action,
	}
	return &c, nil
}
