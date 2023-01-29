package prompt

import (
	"errors"
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
)

type Prompt struct {}

func New() *Prompt {
	return &Prompt{}
}

func (p *Prompt) RenderText(label string) (string, error) {
	response := ""
	err := survey.AskOne(&survey.Input{ Message: label }, &response)

	if err != nil && err == terminal.InterruptErr {
		log.Fatal("interrupted")
	}

	if response == "" {
		return "", errors.New("empty input")
	}

	return response, nil
}

func (p *Prompt) RenderSelect(label string, options []string) int {
	var selectedIndex int
	err := survey.AskOne(&survey.Select{
		Message: label,
		Options: options,
	}, &selectedIndex)

	if err != nil && err == terminal.InterruptErr {
		log.Fatal("interrupted")
	}

	return selectedIndex
}

func (p *Prompt) RenderError(err error) {
	var response bool
	label := err.Error() + "\n" + "would you like to go back to main menu?"

	err = survey.AskOne(&survey.Confirm{ Message: label }, &response)

	if !response {
		os.Exit(0)
	}

	if err != nil && err == terminal.InterruptErr {
		log.Fatal("interrupted")
	}
}

func (p *Prompt) RenderSuccess(text string) {
	var response bool
	label := text + "\n" + "would you like to go back to main menu?"

	err := survey.AskOne(&survey.Confirm{ Message: label }, &response)

	if !response {
		os.Exit(0)
	}

	if err != nil && err == terminal.InterruptErr {
		log.Fatal("interrupted")
	}
}