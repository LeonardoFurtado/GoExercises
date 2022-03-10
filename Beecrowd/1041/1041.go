package main

import "fmt"

func CoordinatePoint(x, y float64) (result string) {
	if x > 0.0 && y < 0.0 {
		result = "Q4"
	} else if x < 0.0 && y < 0.0 {
		result = "Q3"
	} else if x < 0.0 && y > 0.0 {
		result = "Q2"
	} else if x > 0.0 && y > 0.0 {
		result = "Q1"
	} else if x != 0.0 && y == 0.0 {
		result = "Eixo X"
	} else if x == 0.0 && y != 0.0 {
		result = "Eixo Y"
	} else {
		result = "Origem"
	}
	return
}

func main() {
	var x, y float64
	_, _ = fmt.Scanf("%f %f", &x, &y)
	fmt.Println(CoordinatePoint(x, y))
}
