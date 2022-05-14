package main

import (
	"fmt"
)

func main() {
	var amigos [5]string //vetor
	fmt.Println(amigos)
	for i := 0; i < len(amigos); i++ {
		fmt.Print("Digite o nome de um de seus amigos: ")
		fmt.Scanf("%s", &amigos[i])
	}
	fmt.Println(amigos)
}
