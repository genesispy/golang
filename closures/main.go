package main

import (
	"fmt"
)

func main() {
	fmt.Printf("------------- Anotações Closures -------------\n\n")
	fmt.Printf("Basicamente, closure é uma variável onde você pode armazenar não só")
	fmt.Printf("dados primarios, mas pode também armazenar toda uma função dentro de uma var declarada.\n")

	fmt.Printf("Coletando dados!\n")

	var numero1, numero2 int
	var operacao string

	fmt.Print("Digite um numero: \n")
	fmt.Scanf("%d", &numero1)
	fmt.Print("Digite outro numero: \n")
	fmt.Scanf("%d", &numero2)
	fmt.Print("Digite + para somar: \n")
	fmt.Scanf("%s", &operacao)

	fmt.Printf("INICIANDO A CLOSURE!\n")

	var funcaoSoma func(n1 int, n2 int) int //aqui está a closure!
	for operacao != "+" {
		fmt.Print("caracter invalido, tente novamente!\n")
		fmt.Printf("Digite um caracter valido: \n")
		fmt.Scanf("%s", &operacao)
	}
	if operacao == "+" {
		funcaoSoma = func(n1 int, n2 int) int {
			return n1 + n2
		}
	}
	resultado1 := funcaoSoma(numero1, numero2)
	fmt.Printf("%d %s %d = %d", numero1, operacao, numero2, resultado1)
}
