package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var codigo_peca_1, codigo_peca_2, numero_pecas_1, numero_pecas_2 int
	var valor_peca_1, valor_peca_2 float64
	var r = bufio.NewReader(os.Stdin)
	fmt.Fscanf(r, "%d %d %f\n", &codigo_peca_1, &numero_pecas_1, &valor_peca_1)
	fmt.Fscanf(r, "%d %d %f\n", &codigo_peca_2, &numero_pecas_2, &valor_peca_2)

	total := valor_peca_1*float64(numero_pecas_1) + valor_peca_2*float64(numero_pecas_2)

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)
}
