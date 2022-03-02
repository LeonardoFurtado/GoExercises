package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B, C float64
	_, _ = fmt.Scanf("%f %f %f", &A, &B, &C)

	if math.Abs(A-B) < C && C < (A+B) {
		fmt.Printf("Perimetro = %.1f\n", A+B+C)
	} else {
		fmt.Printf("Area = %.1f\n", 0.5*(A+B)*C)
	}
}
