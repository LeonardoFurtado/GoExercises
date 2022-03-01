package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B, C float64
	_, _ = fmt.Scanf("%f %f %f", &A, &B, &C)
	delta := (B * B) - (4 * A * C)
	if A == 0 || delta < 0 {
		fmt.Println("Impossivel calcular")
	} else {
		r1 := (-B + math.Sqrt(delta)) / (2.0 * A)
		r2 := (-B - math.Sqrt(delta)) / (2.0 * A)
		fmt.Printf("R1 = %.5f\n", r1)
		fmt.Printf("R2 = %.5f\n", r2)
	}
}
