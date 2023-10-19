package Tools

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	Sessions "Rain/core/masters/sessions"

	"github.com/fatih/color"
)

type Color struct {
	R, G, B int
}

var _ = reflect.TypeOf(Color{})

func ToRGB(h string) (c Color, err error) {
	switch len(h) {
	case 6:
		_, err = fmt.Sscanf(h, "%02x%02x%02x", &c.R, &c.G, &c.B)
	case 3:
		_, err = fmt.Sscanf(h, "%1x%1x%1x", &c.R, &c.G, &c.B)
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("STATUS") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.YellowString("GRADIENT") + color.WhiteString(":") + color.RedString("FATAL") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.RedString("Failed to load gradient colours") + color.WhiteString("]"))
	}
	return
}

func Bresenham(s, e float64, steps int) []int {
	delta := (e - s) / (float64(steps) - 1)
	colors := []int{int(s)}
	err := 0.0

	for i := 0; i < steps-1; i++ {
		n := float64(colors[i]) + delta
		err = err + (n - float64(int(n)))
		if err >= 0.5 {
			n = n + 1.0
			err = err - 1.0
		}
		colors = append(colors, int(n))
	}
	return colors
}

func Gradient(c1, c2 Color, n int) ([]int, []int, []int) {
	if n < 3 {
		r := []int{c1.R, c2.R}
		g := []int{c1.G, c2.G}
		b := []int{c1.B, c2.B}
		return r, g, b
	}

	R := Bresenham(float64(c1.R), float64(c2.R), n)
	G := Bresenham(float64(c1.G), float64(c2.G), n)
	B := Bresenham(float64(c1.B), float64(c2.B), n)
	return R, G, B
}
func Colorize(text string, r, g, b int) string {
	fg := fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
	return fg + text + "\x1b[0m"
}

func Fade(Input string, Session *Sessions.Session) (string, error) {

	lines := strings.SplitAfter(string(Input), "\n")
	llen := []int{}
	for _, v := range lines {
		llen = append(llen, len(v))
	}

	sort.Ints(llen[:])

	color1, err := ToRGB(Session.ColourOne)
	if err != nil {
		return Input, err
	}

	color2, err := ToRGB(Session.ColourTwo)
	if err != nil {
		return Input, err
	}

	r, g, b := Gradient(color1, color2, llen[len(llen)-1])

	out := []string{}
	for _, line := range lines {
		for i, v := range line {
			out = append(out, Colorize(string(v), r[i], g[i], b[i]))
		}
	}

	return strings.Join(out, ""), nil
}
