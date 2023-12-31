package main

import (
	c "AC3/Contato"
    oa "AC3/OperacoesArray"
	"bufio"
	"fmt"
	"os"
)


func main()  {
	var contatos [5]c.Contato
	var acao int
	var indice int
	var email string

	for true{
		fmt.Println("digite 1 para adicionar, 2 editar email, 3 para excluir, 4 para mostrar contatos, 5 para sair")
		fmt.Scanln(&acao)

		switch acao {
		case 1:
			leitor := bufio.NewReader(os.Stdin)
			fmt.Println("informe seu nome: ")
			nome, _ := leitor.ReadString('\n')

			fmt.Println("informe seu email: ")
			fmt.Scanln(&email)

			contato := c.Contato{
				Nome: nome,
				Email: email,
			}

			oa.AdicionaContato(contato, &contatos)

		case 2:
			oa.ExibirContatos(contatos)
			fmt.Println("escreva o indice do contato que deseja editar o email: ")
			fmt.Scanln(&indice)

			fmt.Println("digite o novo email: ")
			fmt.Scanln(&email)

			oa.EditarEmail(indice, email, &contatos)
		case 3:
			oa.ExcluiContato(&contatos)
		case 4:
			oa.ExibirContatos(contatos)
		case 5:
			break
		default:
			continue
		}
	}
}
