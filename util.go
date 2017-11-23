package main

import "math"

// https://www.desmos.com/calculator/yex502ot5u
func quadGenerator(width, height float64) func(float64) float64 {
	mid := (1 / 2) * width
	stretch := height / (2 * math.Pow(mid, 2))
	return func(x float64) float64 {
		if x >= mid {
			return stretch * math.Pow(x, 2)
		} else {
			return -stretch*math.Pow(x-2*mid, 2) + height
		}
	}
}
