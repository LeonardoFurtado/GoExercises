package main

import "fmt"

func main() {
	I := 1
	J := 60
	for J >= 0 {
		fmt.Printf("I=%d J=%d\n", I, J)
		J -= 5
		I += 3
	}
}
