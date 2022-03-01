package main

import "fmt"

func main() {
	var secondsInput int
	_, _ = fmt.Scanln(&secondsInput)
	hours := secondsInput / 3600
	rest := secondsInput % 3600
	minutes := rest / 60
	rest = rest % 60
	seconds := rest / 1
	fmt.Printf("%d:%d:%d\n", hours, minutes, seconds)
}
