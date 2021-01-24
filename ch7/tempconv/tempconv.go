package tempconv

import (
	"fmt"
)

type Celsius float64

type celsiusFlag struct {Celsius}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", ""
	}
}

