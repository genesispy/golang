package main

import (
	"fmt"
	"strconv"
)

func main() {
	var texto string
	fmt.Print("Digite um numero: ")
	fmt.Scanf("%s", &texto)
	//var numero int
	//numero, _ = strconv.Atoi(texto)

	numero, _ := strconv.ParseInt(texto, 10, 32)
	fmt.Println(numero)
	var conv = float64(numero) //tirando de int e jogando para float64
	fmt.Println(conv)

	testemp := map[string]string{"nome": "Geno", "animal": "gato"}
	fmt.Println(testemp["nome"]) // apenas vendo como funciona mapas em go
}
