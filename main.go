package main

import (
	"github.com/bemsk/parky/app"
	"github.com/bemsk/parky/auxiliaries/palette"
	"github.com/bemsk/parky/auxiliaries/prompt"
	"github.com/bemsk/parky/delivery/cli"
	"github.com/bemsk/parky/parking"
	"github.com/bemsk/parky/repository/memory"
)


func main() {
	repo := memory.New()
	cls := palette.New()
	colors := make([]parking.ReadableColor, len(cls))

	for i, c := range cls {
		colors[i] = c
	}
	
	app := app.New(repo, colors)
	prompter := prompt.New()
	cli := cli.New(app, prompter)

	cli.Run()
}