package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_demo/grpc_python/proto"
)

func main() {
	conn, err := grpc.Dial(":8099", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()

	req := proto.UserRequest{Name: "foo"}

	serviceUserClient := proto.NewServiceUserClient(conn)
	userResp, err := serviceUserClient.GetUserInfo(context.Background(),&req)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(userResp.Name,userResp.Id,userResp.Age,userResp.Hobby)
}
