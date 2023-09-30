package devices

import (
	"fmt"
	"math"
)

type Keyboard struct {
	Name   string
	Layout []KeyboardKey `json:"layout"`
}

type KeyboardKey struct {
	Label string
	Row   int
	Col   int
	X     float64
	Y     float64
	R     float64
	Rx    float64
	Ry    float64
}

const DefaultSize = 65
const DefaultPadding = 5

// / move this to frontend.
func (k *KeyboardKey) Style() string {
	x := k.X * (DefaultSize + DefaultPadding)
	y := k.Y * (DefaultSize + DefaultPadding)
	u := DefaultSize
	h := DefaultSize
	rx := (k.X - math.Max(k.Rx, k.X)) * -(DefaultSize + DefaultPadding)
	ry := (k.Y - math.Max(k.Ry, k.Y)) * -(DefaultSize + DefaultPadding)
	a := k.R

	return fmt.Sprintf(
		"top: %vpx; left: %vpx; width: %vpx; height: %vpx; transform-origin: %vpx %vpx; transform: rotate(%vdeg)",
		y, x, u, h, rx, ry, a)

}
