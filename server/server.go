package main

import(
	pb "github.com/qiansuo1/gRPC_test/proto/todo/v2"
)

type server struct{
	d db

	pb.UnimplementedTodoServiceServer
}