package cli

import (
	"strconv"

	"github.com/bemsk/parky/app"
	"github.com/bemsk/parky/delivery"
)

type ReadableLabel interface {
	Label() string
}

type Prompter interface {
	RenderText(label string) (string, error)
	RenderSelect(label string, options []string) int
	RenderError(err error)
	RenderSuccess(text string)
}

type CLI struct {
	app *app.App
	prompter Prompter
	commands map[string]*Command
}

func New(app *app.App, p Prompter) *CLI {
	cli := &CLI{app, p, make(map[string]*Command)}

	cli.RegisterCommand(0, "park", OnCarEntry)
	cli.RegisterCommand(1, "unpark", OnCarExit)
	cli.RegisterCommand(2, "find slot by registration number", OnSlotQuery)
	cli.RegisterCommand(3, "find registration numbers by color", OnRegistrationNumbersQuery)
	cli.RegisterCommand(4, "find slot numbers by color", OnSlotNumbersQuery)
	cli.RegisterCommand(5, "exit program", OnQuit)

	return cli
}

func (c *CLI) Run() {
	c.PromptParkingCapacity()

	for {
		cmd := c.PromptCommands()

		cmd.Execute(c.app, c.prompter)
	}	
}

func (c *CLI) RegisterCommand(order int, label string, handler func(app *app.App, prompter Prompter)) {
	cmd := &Command{order, label, handler}
	slug := cmd.Slug()

	c.commands[slug] = cmd
}

func (c *CLI) PromptParkingCapacity() {
	capacityStr, err := c.prompter.RenderText("Enter parking capacity:")

	if err != nil {
		c.prompter.RenderError(err)
	}

	capacity, err := strconv.Atoi(capacityStr)

	if err != nil {
		c.prompter.RenderError(delivery.ErrDeliveryInvalidInput)
	}

	c.app.SetParkingCapacity(capacity)
}

func (c *CLI) PromptCommands() *Command {
	cmds := make([]*Command, len(c.commands))

	for _, v := range c.commands {
		cmds[v.Order()] = v
	}

	options := make([]string, len(cmds))

	for i, v := range cmds {
		options[i] = v.Label()
	}

	selectedIndex := c.prompter.RenderSelect("choose an option:", options)

	return cmds[selectedIndex]
}