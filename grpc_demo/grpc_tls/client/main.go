package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc_demo/grpc_tls/proto"
)

func main() {
	cred, err := credentials.NewClientTLSFromFile("../keys/server.pem","xielonglong")
	if err != nil {
		panic(err.Error())
	}
	conn, err := grpc.Dial("localhost:8099",grpc.WithTransportCredentials(cred))
	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()

	var arg1 int64 =2
	var arg2 int64 = 3
	var req = proto.ArgRequest{Arg1: arg1,Arg2: arg2}
	serMathClient := proto.NewServiceMathClient(conn)
	resp, err := serMathClient.GetRes(context.Background(),&req)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(resp.Resp)
}
