/*
	Esse arquivo contêm a função de conexão do client client como o server
*/

package main

import (
	"ListTask/grpc/proto"
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = flag.String("addr", "localhost:9000", "the address to connect to")

func main() {
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := proto.NewTaskListClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	userScreen(client, ctx)

	/*   FUNÇÃO ADDTASK  */
	//_, err = client.AddTask(ctx, &proto.Task{Task: "Treco pra fazer"})

	/*   FUNÇÃO LISTTASKS  */
	//tasks, err := client.ListTasks(ctx, &proto.Void{})

	/*   FUNÇÃO DONETASK  */
	//_, err = client.DoneTask(ctx, &proto.IdTask{Id: 6})

	/*   FUNÇÃO DELETETASK    */
	//_, err = client.DeleteTask(ctx, &proto.Void{})

	/*if err != nil {
		log.Fatalf("could not add task: %v", err)
	}*/

	/*   FUNÇÃO LIST TASK     */
	/*for _, line := range tasks.GetTask() {
		fmt.Printf("%s\n", line)
	}*/
}
