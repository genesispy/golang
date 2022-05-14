package main

import (
	"fmt"
)

func main() {
	fmt.Println("Estou abrindo uma conexão com o banco de dados!")
	defer fmt.Println("Estou fechando a conexão com o banco de dados") /*independente do corpo do codigo, ao
	  final da função essa linha será executada! */
	defer fmt.Println("O ultimo defer é executado primeiro!") //1º defer a ser executado é sempre o primeiro
	executarConsulta()
}

func executarConsulta() {
	fmt.Println("Executando consulta na tabela do database")
}
