package main

import (
	"fmt"
)

func main() {
	var numero int
	fmt.Println("Digite um numero: ")
	fmt.Scanf("%d", &numero)
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d x %d = %d", numero, i, numero*i)
		fmt.Println("\n")
	}
}
