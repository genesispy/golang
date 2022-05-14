package main

import (
	"fmt"
)

func main() {
	var saldo float64
	fmt.Print("Digite seu saldo: ")
	fmt.Scanf("%f", &saldo)
	CalcularRedimento(&saldo) // calcular rendimento de saldo(usando & como ponteiro para ler valor)
	fmt.Printf("Seu saldo é de %.2f\n", saldo)
}

func CalcularRedimento(saldo *float64) { //* é o ponteiro indicando o valor que ocupa a memória da var saldo, por exemplo
	*saldo += *saldo * 0.03 //*saldo = ponteiro chamando o valor na memoria saldo
	//new() cria um ponteiro para alguma área
}
