package main

import "fmt"

func main() {
	var value1 int
	var value2 int

	_, _ = fmt.Scanln(&value1)
	_, _ = fmt.Scanln(&value2)
	PROD := value1 * value2
	fmt.Println("PROD =", PROD)

}
