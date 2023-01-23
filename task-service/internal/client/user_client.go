package client

import (
	"context"
	"fmt"
	"task-service/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserServiceClient struct {
	Client pb.UserServiceClient
}

func InitUserServiceClient(url string) UserServiceClient {
	cc, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := UserServiceClient{
		Client: pb.NewUserServiceClient(cc),
	}

	return c
}

func (s *UserServiceClient) GetUserInfo(ctx context.Context, userId int64) (*pb.GetUserInfoResponse, error) {
	req := &pb.GetUserInfoRequest{
		UserId: userId,
	}
	return s.Client.GetUserInfo(ctx, req)
}
