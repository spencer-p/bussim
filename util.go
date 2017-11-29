package main

import "math"

// https://www.desmos.com/calculator/yex502ot5u
func quadGenerator(width, height float64) func(float64) float64 {
	mid := 0.5 * width
	stretch := height / (2 * math.Pow(mid, 2))
	return func(x float64) float64 {
		if x <= 0 {
			return 0
		} else if x <= mid {
			return stretch * math.Pow(x, 2)
		} else if x < width {
			return -stretch*math.Pow(x-2*mid, 2) + height
		} else {
			return height
		}
	}
}
