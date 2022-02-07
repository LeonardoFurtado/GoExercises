package main

import (
	"fmt"
)

func main() {
	var func_number, worked_hours int
	var money_per_hour, salary float64
	_, _ = fmt.Scanln(&func_number)
	_, _ = fmt.Scanln(&worked_hours)
	_, _ = fmt.Scanln(&money_per_hour)

	salary = money_per_hour * float64(worked_hours)

	fmt.Println("NUMBER =", func_number)
	fmt.Printf("SALARY = U$ %.2f\n", salary)
}
