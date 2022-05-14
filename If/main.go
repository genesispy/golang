package main

import (
	"fmt"
)

func main() {
	/*var idade int
	fmt.Print("Digite sua idade: ")
	fmt.Scanf("%d", &idade)
	if idade >= 18 {
		fmt.Println("\n Maior de idade!")
	} else {
		fmt.Print("\n Menor de idade!")
	}*/

	var numero int
	fmt.Println("Digite um numero: ")
	fmt.Scanf("%d", &numero)
	if numero < 0 {
		fmt.Print("Numero negativo!")
	} else { //também existe else if aqui!!
		fmt.Print("Numero positivo!!")
	}
}
