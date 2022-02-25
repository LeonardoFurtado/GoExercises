package main

import "fmt"

func main() {
	var totalDistancia int
	var totalCombustivel float64
	_, _ = fmt.Scanln(&totalDistancia)
	_, _ = fmt.Scanln(&totalCombustivel)

	result := float64(totalDistancia) / totalCombustivel

	fmt.Printf("%.3f km/l\n", result)
}
