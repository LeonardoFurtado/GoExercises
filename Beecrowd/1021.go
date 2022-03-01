package main

import "fmt"

func main() {
	var valor float64
	var res int
	notaOrMoeda := "nota"
	_, _ = fmt.Scanln(&valor)
	res = int((valor * 100.0) + 0.5)
	var cedulas = [12]int{10000, 5000, 2000, 1000, 500, 200, 100, 50, 25, 10, 5, 1}
	fmt.Println("NOTAS:")
	for i := 0; i < len(cedulas); i++ {
		if i == 6 {
			fmt.Println("MOEDAS:")
			notaOrMoeda = "moeda"
		}
		quantidade := res / cedulas[i]
		res %= cedulas[i]
		fmt.Printf("%d %s(s) de R$ %.2f\n", quantidade, notaOrMoeda, float64(cedulas[i])/100.0)
	}
}
