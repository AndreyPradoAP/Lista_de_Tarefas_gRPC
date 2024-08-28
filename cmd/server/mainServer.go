/*
	Esse arquivo contêm todo os códigos pra manter o server rodando
*/

package main

import (
	"ListTask/grpc/proto"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

var port = flag.Int("port", 9000, "The server port")

type Server struct {
	proto.UnimplementedTaskListServer
}

// Função para adicionar tarefa ao arquivo
// Ele aponta para a função serverAddTask() presente no arquivo controllerServer.go
func (server *Server) AddTask(ctx context.Context, task *proto.Task) (*proto.Void, error) {
	println("client adding data")
	return &proto.Void{}, serverAddTask(task.Task)
}

// Função para marcar uma tarefa como concluida
// Ele aponta para a função serverDoneTask() presente no arquivo controllerServer.go
// Arrumar essa disgraça que tá apagando tudo meu arquivo
func (server *Server) DoneTask(ctx context.Context, idTask *proto.IdTask) (*proto.Void, error) {
	println("client changing task as completed")
	return &proto.Void{}, serverDoneTask(strconv.Itoa(int(idTask.GetId())))
}

// Função para listar as tarefas
// Ele aponta para a função serverListTasks() presente no arquivo controllerServer.go
func (server *Server) ListTasks(ctx context.Context, void *proto.Void) (*proto.Tasks, error) {
	println("client requesting data")
	str, err := serverListTasks()

	println("requisition completed")
	return &proto.Tasks{Task: str}, err
}

// Função para deletar todas as tarefas
// Ele aponta para a função deleteData() presente no arquivo controllerServer.go
func (server *Server) DeleteTask(ctx context.Context, void *proto.Void) (*proto.Void, error) {
	println("client deleting all data")
	err := serverDeleteData()
	return &proto.Void{}, err
}

// Deixa o servidor atento às requisições do cliente
func main() {
	println("gRPC server running")

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	proto.RegisterTaskListServer(server, &Server{})
	log.Printf("server listen at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
