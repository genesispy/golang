package main

import (
	"fmt" //importando biblioteca fmt
)

func main() {
	amigos := make([]string, 3) //iniciando um slice de 3 espaços
	nome := ""
	i := 0
	for nome != "q" { //for como while e definindo 'q' como "break"
		fmt.Printf("Digite um nome ou 'q' para encerrar: ")
		fmt.Scanf("%s", &nome)
		if i > 3 { //se as entradas registradas no slices passarem do limite
			amigos = append(amigos, nome) //inserindo entradas a mais no slice
		}
	}
	fmt.Printf("Você tem %d amigos", len(amigos)) //printando a quantidade de entradas
}
