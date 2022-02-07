package main

import "fmt"


func main(){
	var A int
	var B int
	_, _ = fmt.Scanln(&A)
	_, _ = fmt.Scanln(&B)
	SOMA := A + B
	fmt.Println("SOMA =", SOMA)
}


