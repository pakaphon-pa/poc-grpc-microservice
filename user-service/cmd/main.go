package main

import (
	"context"
	"log"
	"net"
	"user-service-grpc/pb"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func (s *Server) GetUserInfo(ctx context.Context, in *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	return &pb.GetUserInfoResponse{Id: 1, Name: "Test", Email: "Test@Test.com"}, nil
}

func main() {
	println("Running RPC server...")

	lis, err := net.Listen("tcp", ":8082")

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
