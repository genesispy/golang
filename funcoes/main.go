package main

import (
	"fmt"
)

func main() {
	var numero1, numero2 int
	var operacao = "+"
	fmt.Println("Digite o primeiro numero: ")
	fmt.Scanf("%d", &numero1)
	fmt.Printf("Digite a operação (+ - * /): ")
	fmt.Scanf("%s", &operacao)
	fmt.Println("Digite o segundo numero: ")
	fmt.Scanf("%d", &numero2)
	if operacao == "+" {
		resultado := soma(numero1, numero2)
		fmt.Println(resultado)
	} else if operacao == "-" {
		resultado := subtrair(numero1, numero2)
		fmt.Println(resultado)
	} else if operacao == "*" {
		resultado := multiplicar(numero1, numero2)
		fmt.Println(resultado)
	} else if operacao == "/" {
		resultado := dividir(numero1, numero2)
		fmt.Println(resultado)

	}

}

func soma(n1 int, n2 int) int {
	return (n1 + n2)
}

func subtrair(n1 int, n2 int) int {
	return (n1 - n2)
}

func multiplicar(n1 int, n2 int) int {
	return (n1 * n2)
}

func dividir(n1 int, n2 int) int {
	return (n1 / n2)
}
