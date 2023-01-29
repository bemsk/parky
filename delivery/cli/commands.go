package cli

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bemsk/parky/app"
)

type Command struct {
	order int
	label string
	handler func(app *app.App, prompter Prompter)
}

func (c *Command) Order() int {
	return c.order
}

func (c *Command) Label() string {
	return c.label
}

func (c *Command) Slug() string {
	return strings.Join(strings.SplitAfter(c.label, " "), "-")
}

func (c *Command) Execute(app *app.App, p Prompter) {
	c.handler(app, p)
}

func OnCarEntry(app *app.App, p Prompter) {
	registrationNumber, err := p.RenderText("Registration Number")

	if err != nil {
		p.RenderError(err)
		return
	}

	colors := app.ListAvailableColors()
	options := make([]string, len(colors))
	
	for i, c := range colors {
		options[i] =  c.ReadableName()
	}

	selectedIndex := p.RenderSelect("Select color", options)

	selectedColor := colors[selectedIndex]
	ticket, err := app.OrderTicket(registrationNumber, selectedColor)

	if err != nil {
		p.RenderError(err)
		return
	}

	p.RenderSuccess("Go to slot number: " + strconv.Itoa(ticket.SlotNumber))
}

func OnCarExit(app *app.App, p Prompter) {
	registrationNumber, err := p.RenderText("Registration Number")

	if err != nil {
		p.RenderError(err)
		return
	}

	err = app.RetractTicket(registrationNumber)

	if err != nil {
		p.RenderError(err)
		return
	}

	p.RenderSuccess("Exit success")
}

func OnSlotQuery(app *app.App, p Prompter) {
	registrationNumber, err := p.RenderText("Registration Number")

	if err != nil {
		p.RenderError(err)
		return
	}

	sn, err := app.SlotNumber(registrationNumber)

	if err != nil {
		p.RenderError(err)
		return
	}

	p.RenderSuccess("Car is available on slot number " + strconv.Itoa(sn))
}

func OnRegistrationNumbersQuery(app *app.App, p Prompter) {
	colors := app.ListAvailableColors()
	options := make([]string, len(colors))

	for i, c := range colors {
		options[i] =  c.ReadableName()
	}

	selectedIndex := p.RenderSelect("Select color", options)

	selectedColor := colors[selectedIndex]

	rns := app.RegistrationNumbersByColor(selectedColor)
	
	if (len(rns) > 0) {
		p.RenderSuccess("cars with " + selectedColor.ReadableName() + " color:" + strings.Join(rns, ", "))
	} else {
		p.RenderSuccess("currently there are no car with such color in the parking lot")
	}
}

func OnSlotNumbersQuery(app *app.App, p Prompter) {
	colors := app.ListAvailableColors()
	options := make([]string, len(colors))

	for i, c := range colors {
		options[i] =  c.ReadableName()
	}

	selectedIndex := p.RenderSelect("select color", options)

	selectedColor := colors[selectedIndex]

	sns := app.SlotNumbersByColor(selectedColor)
	slots := make([]string, len(sns))

	for i, v := range sns {
		slots[i] = strconv.Itoa(v)
	}

	if (len(sns) > 0) {
		p.RenderSuccess("slots with " + selectedColor.ReadableName() + " color car in it:" + strings.Join(slots, ", "))
	} else {
		p.RenderSuccess("currently there are no slot occupied by " + selectedColor.ReadableName() + " color car in the parking lot")
	}
}

func OnQuit(app *app.App, p Prompter) {
	fmt.Println("Goodbye!")
	os.Exit(0)
}