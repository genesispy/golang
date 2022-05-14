package main

import (
	"fmt"
)

func main() {
	var numero int
	fmt.Println("Digite um numero: ")
	fmt.Scanf("%d", &numero)
	i := 0
	for i <= 10 {
		fmt.Printf("%d x %d = %d \n", numero, i, numero*i)
		i++
	}
}
