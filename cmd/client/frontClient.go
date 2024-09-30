/*
	Esse arquivo contêm os códigos que serão vistos e interagirão com o usuário, o qual acessará o server
*/

package main

import (
	"ListTask/grpc/proto"
	"bufio"
	"context"
	"fmt"
	"os"
)

/*func clientAddTask() {

}*/

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
			// Por algum motivo, a função funciona nos primeiros segundos de rodagem, depois para de enviar as requisições por conta do seguinte erro:
			// rpc error: code = DeadlineExceeded desc = context deadline exceeded (MESMO ESTANDO NO LOCALHOST)
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Printf("\nAdição de tarefa escolhida\n")
			fmt.Printf("\tDigite a descrição da nova tarefa\n\t:>")
			scanner.Scan()
			descTask := scanner.Text()
			_, err := client.AddTask(ctx, &proto.Task{Task: descTask})
			if err != nil {
				errorMessage(err)
				break
			}

			fmt.Printf("Tarefa adicionada à lista com exito!\n\n")

		case 2:
			var doneTaskInd int
			fmt.Printf("\nConclusão de tarefa escolhida\n")
			fmt.Printf("\tDigite o índice da tarefa a ser concluída\n\t:>")
			fmt.Scanf("%d", &doneTaskInd)
			_, err := client.DoneTask(ctx, &proto.IdTask{Id: int32(doneTaskInd)})
			if err != nil {
				errorMessage(err)
				break
			}

			fmt.Printf("Tarefa %d concluída com sucesso!\n\n", doneTaskInd)

		case 3:
			fmt.Printf("\nListagem de tarefa escolhida\n")
			tasks, err := client.ListTasks(ctx, &proto.Void{})

			if err != nil {
				errorMessage(err)
				break
			}

			for _, line := range tasks.GetTask() {
				fmt.Printf("\t%s\n", line)
			}

			fmt.Printf("\nListagem realizada com sucesso!\n\n")
		case 4:
			fmt.Printf("Exclusão de tarefas escolhida\n")
			_, err := client.DeleteTask(ctx, &proto.Void{})

			if err != nil {
				errorMessage(err)
				break
			}

			fmt.Printf("\nExclusão realizada com sucesso!\n\n")

		case 5:
			fmt.Println("Programa encerrado!")
			return
		default:
			fmt.Printf("Opção inválida, digite novamente!\n")
		}
	}
}
