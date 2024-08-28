/*
	Esse arquivo contêm os códigos que serão vistos e interagirão com o usuário, o qual acessará o server
*/

package main

import (
	"ListTask/grpc/proto"
	"context"
	"fmt"
)

func clientAddTask() {

}

func errorMessage(err error) {
	fmt.Printf("Erro ao executar tarefa: %s\n", err)
}

func userScreen(client proto.TaskListClient, ctx context.Context) {
	var choice int

	for choice != 5 {
		choice = 0
		fmt.Printf("Escolha uma das opções abaixo:\n")
		fmt.Printf("\t1 - Adicionar uma nova tarefa\n")
		fmt.Printf("\t2 - Concluir uma tarefa\n")
		fmt.Printf("\t3 - Vizualizar tarefas\n")
		fmt.Printf("\t4 - Apagar dados\n")
		fmt.Printf("\t5 - Sair\n\n")

		fmt.Printf("Digite o número da Escolha\n:> ")
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1:
			fmt.Printf("Adição de tarefa escolhida\n")
			_, err := client.AddTask(ctx, &proto.Task{Task: "Treco pra fazer"})
			if err != nil {
				errorMessage(err)
			} else {
				fmt.Printf("Tarefa adicionada à lista com exito!\n")
			}

		case 2:
			//doneTask()
			fmt.Println("case 2")
		case 3:
			fmt.Println("case 3")
		case 4:
			fmt.Println("case 4")
		case 5:
			fmt.Println("Programa encerrado!")
			return
		default:
			fmt.Printf("Opção inválida, digite novamente!\n")
		}
	}
}