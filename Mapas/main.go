package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	mapaCursos := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	curso := ""
	for curso != "q" {
		var cargaHoraria int
		fmt.Println("Digite um curso ou 'q' para sair: ")
		scanner.Scan()
		curso = scanner.Text()
		if curso != "q" {
			fmt.Println("Digite a carga horaria do curso: ")
			fmt.Scanf("%d", &cargaHoraria)
			mapaCursos[curso] = cargaHoraria
		}
	}
	fmt.Println("Cursos registrados: ")
	for nomeCurso, cargaHoraria := range mapaCursos {
		fmt.Printf(" - %s: %dh \n", nomeCurso, cargaHoraria)
	}
}
