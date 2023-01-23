package main

import (
	"api-gateway/pb"
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cc, err := grpc.Dial("task-service:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	taskService := pb.NewTaskServiceClient(cc)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		fmt.Println("Request api gateway")
		req := &pb.CreateTaskRequest{Name: 1, Description: 1, UserId: 1}
		res, _ := taskService.CreateTask(context.Background(), req)
		return c.JSON(http.StatusOK, res)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
