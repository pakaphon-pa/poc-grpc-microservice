package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"task-service/internal/client"
	"task-service/pb"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedTaskServiceServer
	UserService client.UserServiceClient
}

func (s *Server) CreateTask(ctx context.Context, in *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	user, _ := s.UserService.GetUserInfo(ctx, in.UserId)
	fmt.Println("UserId: " + strconv.FormatInt(user.Id, 10) + "Email: " + user.Email)
	return &pb.CreateTaskResponse{Status: 200, Id: 1, Error: ""}, nil
}

func main() {
	println("Running RPC server...")

	lis, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}
	userService := client.InitUserServiceClient("user-service:8082")

	grpcServer := grpc.NewServer()

	s := Server{
		UserService: userService,
	}

	pb.RegisterTaskServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
