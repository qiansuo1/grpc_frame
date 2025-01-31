package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	pb "github.com/qiansuo1/gRPC_test/proto/todo/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {

	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalln("usage: client [IP_ADDR]")
	}
	addr := args[0]

	cerds, err := credentials.NewClientTLSFromFile("./certs/ca_cert.pem", "x.test.example.com")
	 if err != nil{
		log.Fatalf("failed to load credentials: %v", err)
	 }

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(cerds),
		//grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(unaryAuthInterceptor),
		grpc.WithStreamInterceptor(streamAuthInterceptor),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
	}
	
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer func(conn *grpc.ClientConn) {
		if err := conn.Close(); err != nil {
			log.Fatalf("unexpected error: %v", err)
		}
	}(conn)

	c := pb.NewTodoServiceClient(conn)

	fm, err := fieldmaskpb.New(&pb.Task{}, "id")

	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
	fmt.Println("--------ADD--------")
	dueDate := time.Now().Add(5 * time.Second)
	id1 := addTask(c, "This is a task", dueDate)
	id2 := addTask(c, "This is another task", dueDate)
	id3 := addTask(c, "And yet another task", dueDate)
	fmt.Println("-------------------")

	fmt.Println("--------LIST-------")
	printTasks(c, fm)
	fmt.Println("-------------------")

	fmt.Println("-------UPDATE------")
	updateTasks(c, []*pb.UpdateTasksRequest{
		{Id: id1, Description: "A better name for the task"},
		{Id: id2, DueDate: timestamppb.New(dueDate.Add(5 * time.Hour))},
		{Id: id3, Done: true},
	}...)
	printTasks(c, fm)
	fmt.Println("-------------------")

	fmt.Println("-------DELETE------")
	deleteTasks(c, []*pb.DeleteTasksRequest{
		{Id: id1},
		{Id: id2},
		{Id: id3},
	}...)

	printTasks(c, fm)
	fmt.Println("-------------------")

}

func addTask(c pb.TodoServiceClient, description string, dueDate time.Time) uint64 {
	req := &pb.AddTaskRequest{
		Description: description,
		DueDate:     timestamppb.New(dueDate),
	}
	res, err := c.AddTask(context.Background(), req)
	if err != nil {
		if s, ok := status.FromError(err); ok {
			switch s.Code() {
			case codes.InvalidArgument, codes.Internal:
				log.Fatalf("%s: %s", s.Code(), s.Message())
			default:
				log.Fatal(s)
			}

		} else {
			panic(err)
		}
	}
	fmt.Printf("added task: %d\n", res.Id)
	return res.Id

}

func printTasks(c pb.TodoServiceClient, fm *fieldmaskpb.FieldMask) {
	req := &pb.ListTasksRequest{
		Mask: fm,
	}
	strem, err := c.ListTasks(context.Background(), req)
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}

	for {
		res, err := strem.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("unexpected error: %v", err)
		}
		fmt.Println(res.Task.String(), "overdue: ", res.Overdue)
	}
}

func updateTasks(c pb.TodoServiceClient, reqs ...*pb.UpdateTasksRequest) {
	stream, err := c.UpdateTasks(context.Background())
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}

	for _, req := range reqs {
		err := stream.Send(req)

		if err != nil {
			log.Fatalf("unexpected error: %v", err)
		}

		if req != nil {
			fmt.Printf("updated task with id: %d\n", req.Id)
		}
	}

	if _, err = stream.CloseAndRecv(); err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
}

func deleteTasks(c pb.TodoServiceClient, reqs ...*pb.DeleteTasksRequest) {
	stream, err := c.DeleteTasks(context.Background())

	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			_, err := stream.Recv()

			if err == io.EOF {
				close(waitc)
				break
			}
			if err != nil {
				log.Fatalf("error while receiving: %v\n", err)
			}

			log.Println("deleted tasks")
		}
	}()

	for _, req := range reqs {
		if err := stream.Send(req); err != nil {
			return
		}
	}
	if err := stream.CloseSend(); err != nil {
		return
	}

	<-waitc
}
