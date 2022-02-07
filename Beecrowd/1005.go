package main

import (
	"fmt"
)

func main() {
	var A, B, C, D, DIFERENCA int
	_, _ = fmt.Scanln(&A)
	_, _ = fmt.Scanln(&B)
	_, _ = fmt.Scanln(&C)
	_, _ = fmt.Scanln(&D)
	DIFERENCA = A*B - C*D
	fmt.Println("DIFERENCA =", DIFERENCA)
}
