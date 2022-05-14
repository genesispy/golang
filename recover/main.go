package main

import (
	"fmt"
)

func main() {
	fmt.Println("Bem vindo!!")
	defer func() {
		if ex := recover(); ex != nil { //recover recupera a ultima situação de erro ou nil caso não houver erro
			fmt.Printf("\nOps, um erro foi encontrado: %s", ex)
		}
		fmt.Printf("\nObrigado por usar nosso software!!")
	}() //garanta que o defer seja chamado apos o main fechar
	var num int
	fmt.Print("\nDigite um numero maior que 10: ")
	fmt.Scanf("%d", &num)
	if num <= 10 {
		panic("\nNúmero invalido!!") //exception
	} else {
		fmt.Print("Ótimo, você conseguiu!!")
	}
}
