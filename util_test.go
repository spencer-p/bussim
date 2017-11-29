package main

import (
	"fmt"
)

func SimpleQuad() {
	q := quadGenerator(10, 10)
	for i := 0; i <= 10; i++ {
		fmt.Println(q(float64(i)))
	}
	// Output:
	// 0
	// 0.2
	// 0.8
	// 1.8
	// 3.2
	// 5
	// 6.8
	// 8.2
	// 9.2
	// 9.8
	// 10
}
