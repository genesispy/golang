package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//ARQUIVO indica o local onde os contatos serao salvos
const ARQUIVO string = "agenda.txt"

//interface que especifica como um item pode ser convertido para string
type ConversivelParaString interface {
	toString() string
}

//Contato: estrutura que representa um contato
type Contato struct {
	nome         string
	formaContato string
	valorContato string
}

func (contato *Contato) toString() string {
	return fmt.Sprintf("%s|%s|%s \n", contato.nome, contato.formaContato, contato.valorContato)
}

//Gerenciador contatos : Responsavel por gerenciar os cntatos
type GerenciadorContatos struct{}

func (gerenciador *GerenciadorContatos) carregarContatos() ([]Contato, error) {
	contatos := make([]Contato, 0)
	if _, e := os.Stat(ARQUIVO); !os.IsNotExist(e) {

		arquivo, err := os.Open(ARQUIVO)
		if err != nil {
			return contatos, err
		}
		defer arquivo.Close()
		scanner := bufio.NewScanner(arquivo)
		for scanner.Scan() {
			linhaContato := scanner.Text()
			infoContato := strings.Split(linhaContato, "|")
			contatos = append(contatos, Contato{nome: infoContato[0], formaContato: infoContato[1], valorContato: infoContato[2]})

		}
	}
	return contatos, nil
}

func (geranciador *GerenciadorContatos) salvarContato(contato ConversivelParaString) error {
	arquivo, err := os.OpenFile(ARQUIVO, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer arquivo.Close()
	if _, e := arquivo.WriteString(contato.toString()); e != nil {
		return e
	}
	return nil
}

func main() {
	gerenciador := GerenciadorContatos{}
	opcao := 0
	for true {
		fmt.Println("O que voce deseja fazer? ")
		fmt.Println(" 1 - Listar Contatos")
		fmt.Println(" 2 - Criar Contato")
		fmt.Println(" 3 - Sair")
		switch opcao {
		case 1:
			listarContatos(&gerenciador)
		case 2:
			criarNovoContato(&gerenciador)
		}
		if opcao == 3 {
			break
		}
	}
}

func listarContatos(gerenciador *GerenciadorContatos) {
	contatos, err := gerenciador.carregarContatos()
	if err != nil {
		fmt.Printf("Não foi possivel carregar os contatos %s\n", err)

	} else {
		fmt.Println("Lista de contatos: ")
		for _, contato := range contatos {
			fmt.Printf("- %s, %s: %s \n", contato.nome, contato.formaContato, contato.valorContato)
		}
	}
}

func criarNovoContato(gerenciador *GerenciadorContatos) {
	novoContato := Contato{}
	fmt.Print("Nome do contato: \n")
	fmt.Scanf("%s", &novoContato.nome)
	fmt.Print("Tipo do contato: \n")
	fmt.Scanf("%s", &novoContato.formaContato)
	fmt.Print("Contato: \n")
	fmt.Scanf("%s", &novoContato.valorContato)
	err := gerenciador.salvarContato(&novoContato)
	if err != nil {
		fmt.Printf("Erro ao salvar contato: %s \n", err)
	}
}
