package main

import (
	"fmt"
)

func main() {
	var result, r float64
	_, _ = fmt.Scanln(&r)
	pi := 3.14159
	result = 4.0 / 3.0 * pi * r * r * r
	fmt.Printf("VOLUME = %.3f", result)
}
