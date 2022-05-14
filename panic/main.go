package main

import (
	"fmt"
)

func main() {
	//inicio
	var numero int
	fmt.Print("Bem vindo!")
	defer fmt.Println("\nObrigado por usar nosso software!! :D") //saída do sistema
	//

	//corpo
	fmt.Print("\nDigite um número maior que 10: ")
	fmt.Scanf("%d", &numero)
	//

	//analisando o numero digitado
	if numero <= 10 {
		panic("\nNúmero invalido!!!!!!!")
	} else {
		fmt.Print("\nCorreto!! :)")
	}

}
