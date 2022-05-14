package main

import (
	"fmt"
)

func main() {
	var numero int
	fmt.Println("Digite um numero: ")
	fmt.Scanf("%d", &numero)
	switch numero {
	case 1: //caso for 1
		fmt.Println("Esse é o numero do lobo!")
	case 2: //caso for 2
		fmt.Println("Esse é o numero da vaca!")
	default: //tipo o else
		fmt.Println("Caracter invalido")
	}
	switch numero {
	case 1:
		fmt.Println("Animal Feroz!")
	case 2:
		fmt.Println("Olha o carro de leite!")
	default:
		fmt.Println("Caracter invalido!")
	}
}
