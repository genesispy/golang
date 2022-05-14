package main

import (
	"fmt"
)

func main() {
	var numero1 int
	var numero2 int
	fmt.Print("Digite o primeiro numero: ")
	fmt.Scanln(&numero1) // joga na variavel numero1 "&"
	fmt.Print("Digite o segundo numero: ")
	fmt.Scanln(&numero2)

	fmt.Printf("\n%d + %d = %d \n", numero1, numero2, numero1+numero2)
	fmt.Printf("\n%d - %d = %d \n", numero1, numero2, numero1-numero2)
	fmt.Printf("\n%d x %d = %d \n", numero1, numero2, numero1*numero2)
	fmt.Printf("\n%d / %d = %d \n", numero1, numero2, numero1/numero2)
	fmt.Printf("\n%d %% %d = %d \n", numero1, numero2, numero1%numero2)

}
