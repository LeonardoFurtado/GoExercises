package main

import "fmt"

func main() {
	var input float64
	_, _ = fmt.Scanln(&input)
	if input < 0 || input > 100 {
		fmt.Println("Fora de intervalo")
	} else if input >= 0 && input <= 25 {
		fmt.Println("Intervalo [0,25]")
	} else if input <= 50 {
		fmt.Println("Intervalo (25,50]")
	} else if input <= 75 {
		fmt.Println("Intervalo (50,75]")
	} else if input <= 100 {
		fmt.Println("Intervalo (75,100]")
	}
}
