package main

import "fmt"

func main() {
	var X, Y int
	var vetor = [5]float64{4.0, 4.5, 5.0, 2.0, 1.5}
	_, _ = fmt.Scanf("%d %d", &X, &Y)
	fmt.Printf("Total: R$ %.2f\n", vetor[X-1]*float64(Y))
}
