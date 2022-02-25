package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int
	_, _ = fmt.Scanf("%d", &a)
	_, _ = fmt.Scanf("%d", &b)
	_, _ = fmt.Scanf("%d", &c)
	maiorab := (a + b + int(math.Abs(float64(a-b)))) / 2
	maiorall := (maiorab + c + int(math.Abs(float64(maiorab-c)))) / 2
	fmt.Println(maiorall, "eh o maior")
}
