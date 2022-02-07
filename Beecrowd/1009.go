package main

import (
	"fmt"
)

func main() {
	var nomeVendedor string
	var salario float64
	var totalVendas float64
	comissao := 0.15

	_, _ = fmt.Scanln(&nomeVendedor)
	_, _ = fmt.Scanln(&salario)
	_, _ = fmt.Scanln(&totalVendas)

	total := totalVendas*comissao + salario

	fmt.Printf("TOTAL = R$ %.2f\n", total)
}
