package main

import "fmt"

func main() {
	var spentTime, averageSpeed int
	_, _ = fmt.Scanln(&spentTime)
	_, _ = fmt.Scanln(&averageSpeed)
	distance := averageSpeed * spentTime
	kmPerLiter := 12.0
	litersNeeded := float64(distance) / kmPerLiter
	fmt.Printf("%.3f\n", litersNeeded)
}
