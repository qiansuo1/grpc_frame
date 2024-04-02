package main

import (
	"context"
	"io"
	"log"
	"slices"
	"time"

	pb "github.com/qiansuo1/gRPC_test/proto/todo/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func Filter(msg proto.Message ,mask *fieldmaskpb.FieldMask){
	if mask == nil || len(mask.Paths) == 0{
		return
	}

	rft := msg.ProtoReflect()

	rft.Range(func(fd protoreflect.FieldDescriptor,_ protoreflect.Value) bool {
		if !slices.Contains(mask.Paths,string(fd.Name())){
			rft.Clear(fd)
		}
		return true
	})

}

func (s *server) AddTask(_ context.Context,in *pb.AddTaskRequest)(*pb.AddTaskResponse,error){
	
	// if len(in.Description) == 0{
	// 	return nil,status.Error(
	// 		codes.InvalidArgument,
	// 		"expected a task description,but got an empty string",
	// 	)
	// }

	// if in.DueDate.AsTime().Before(time.Now().UTC()){
	// 	return nil,status.Error(
	// 		codes.InvalidArgument,
	// 		"expected a task due_time that in the future",
	// 	)
	// }

	if err := in.Validate(); err != nil{
		return nil,err
	}

	id, err := s.d.addTask(in.Description, in.DueDate.AsTime())
	if err != nil{
		return nil,status.Errorf(
			codes.Internal,
			"unexpected err:%s",
			err.Error(),
		)
	}
	return &pb.AddTaskResponse{Id: id}, nil
}

func(s *server) ListTask(req *pb.ListTasksRequest,stream pb.TodoService_ListTasksServer) error{
	return s.d.getTasks(func(t interface{}) error{
		ctx := stream.Context()
		select{
		case <- ctx.Done():
			switch ctx.Err() {
			case context.Canceled:
				log.Printf("request canceled: %s", ctx.Err())
			case context.DeadlineExceeded:
				log.Printf("request deadline exceeded: %s", ctx.Err())
			}
			return ctx.Err()
		default:
		}

		task := t.(*pb.Task)
		Filter(task,req.Mask)
		overdue := task.DueDate != nil && !task.Done && task.DueDate.AsTime().Before(time.Now().UTC())
		err := stream.Send(&pb.ListTasksResponse{
			Task:    task,
			Overdue: overdue,
		})
		return err
	} )


}

func(s *server)UpdateTask(stream pb.TodoService_UpdateTasksServer)error{
	for {
		req,err:= stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.UpdateTasksResponse{})
		}

		if err != nil {
			return err
		}

		s.d.updateTask(
			req.Id,
			req.Description,
			req.DueDate.AsTime(),
			req.Done,
		)
	}

}

func (s *server) DeleteTasks(stream pb.TodoService_DeleteTasksServer) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		s.d.deleteTask(req.Id)
		stream.Send(&pb.DeleteTasksResponse{})
	}
}