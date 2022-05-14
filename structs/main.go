package main

import (
	"errors"
	"fmt"
)

type Carro struct {
	modelo     string
	marca      string
	velocidade float32
}

func (carro *Carro) acelerar() error {
	if carro.velocidade < 15 {
		carro.velocidade += 5
		return nil
	} else {
		return errors.New("carro ja atingiu velocidade máxima")
	}
}

func (carro *Carro) frear() error {
	if carro.velocidade > 0 {
		carro.velocidade -= 5
		return nil
	} else {
		return errors.New("carro ja esta parado")
	}
}

func main() {
	defer fmt.Println("Obrigado por usar nosso software!")
	defer fmt.Println("Encerrando.....")
	carro := Carro{modelo: "Palio", marca: "Fiat"}
	opcao := 0
	for opcao != 3 {
		fmt.Println("O que deseja fazer? ")
		fmt.Println("1- acelerar\n 2- frear\n 3- sair")
		fmt.Scanf("%d", &opcao)
		var err error = nil
		switch opcao {
		case 1:
			err = carro.acelerar()
		case 2:
			err = carro.frear()
		}
		if err != nil {
			fmt.Printf("** Não foi possivel executar a ação %d ** \n", err)
		} else if opcao != 3 {
			fmt.Printf("o carro %s da marca %s está a %.2f km/h \n", carro.modelo, carro.marca, carro.velocidade)
		}
	}
}
