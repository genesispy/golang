package main

import (
	"fmt"
	"time"
)

func main() {
	var limite int
	fmt.Print("Informe um limite: ")
	fmt.Scanf("%d", &limite)
	go conteAte(limite)
	for i := 0; i <= limite*10; i++ {
		fmt.Printf("- [main] O número é %d \n", i)
	}
	time.Sleep(10 * time.Second)
}

func conteAte(limite int) {
	for i := 0; i <= limite*20; i++ {
		fmt.Printf(" - [conteAte] O numero é %d\n", i)
	}
}
