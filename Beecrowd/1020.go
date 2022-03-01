package main

import "fmt"

func main() {
	var daysInput int
	_, _ = fmt.Scanln(&daysInput)
	years := daysInput / 365
	rest := daysInput % 365
	months := rest / 30
	rest = rest % 30
	days := rest / 1
	fmt.Println(years, "ano(s)")
	fmt.Println(months, "mes(es)")
	fmt.Println(days, "dia(s)")
}
