// package color implements 16 named colors defined in the HTML 4.01 specification
package palette

import (
	_ "embed"
	"encoding/json"
)

//go:embed palette.json
var paletteByte []byte

type RGB [3]int

type List struct {
	Colors []*Named `json:"colors"`
}

type Named struct {
	Name string `json:"name"`
	RGB *RGB `json:"rgb"`
	TextColor *RGB `json:"text_color_rgb"`
}

func (n *Named) ReadableName() string {
	return n.Name
}

func New() []*Named {
	var list List
	json.Unmarshal(paletteByte, &list)

	return list.Colors
}