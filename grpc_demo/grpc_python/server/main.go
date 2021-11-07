package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_demo/grpc_python/proto"
	"net"
)

type UserMessage struct {}

func (um *UserMessage)GetUserInfo(ctx context.Context, request *proto.UserRequest) (*proto.UserResponse, error)  {
	name := request.Name
	response := new(proto.UserResponse)
	fmt.Println("请求name",name)

	if name == "foo" {
		response = &proto.UserResponse{
			Name: name,
			Id: 1,
			Age: 18,
			Hobby: []string{"test01","test02"},
		}
		return response,nil
	}
	return response,fmt.Errorf("unknow user name=%s",name)
}

func main() {

	userMessage := new(UserMessage)

	server := grpc.NewServer()
	proto.RegisterServiceUserServer(server,userMessage)

	lis, err := net.Listen("tcp","localhost:8099")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(lis)
}