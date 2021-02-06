package colorpoint

import "image/color"

type Point struct{ X, Y float64 }

type ColorPoint struct {
	Point
	Color color.RGBA
}
