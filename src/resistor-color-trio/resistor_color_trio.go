package resistorcolortrio

import (
	"fmt"
	"math"
)

var kiloOhms = 1_000
var megaOhms = 1_000_000
var gigaOhms = 1_000_000_000

var colorMap = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	value := getValue(colors)

	switch {
	case value >= gigaOhms:
		return fmt.Sprintf("%d %s", value/gigaOhms, "gigaohms")
	case value >= megaOhms:
		return fmt.Sprintf("%d %s", value/megaOhms, "megaohms")
	case value >= kiloOhms:
		return fmt.Sprintf("%d %s", value/kiloOhms, "kiloohms")
	case value >= 0:
		return fmt.Sprintf("%d %s", value, "ohms")
	}

	return ""
}

func getValue(colors []string) int {
	return (colorMap[colors[0]]*10 + colorMap[colors[1]]) *
		int(math.Pow10(colorMap[colors[2]]))
}
