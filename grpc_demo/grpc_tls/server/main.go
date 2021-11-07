package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc_demo/grpc_tls/proto"
	"net"
)

type MathMessage struct {}

func (m *MathMessage) GetRes(ctx context.Context, req *proto.ArgRequest) (*proto.MathResp, error)  {
	var resp = new(proto.MathResp)
	resp.Resp = req.Arg1 + req.Arg2
	fmt.Println(resp.Resp)
	return resp,nil
}

func main() {

	var mathMessage = new(MathMessage)
	cred,err := credentials.NewServerTLSFromFile("../keys/server.pem","../keys/server.key")
	if err != nil {
		panic(err.Error())
	}

	server := grpc.NewServer(grpc.Creds(cred))
	proto.RegisterServiceMathServer(server,mathMessage)

	lis, err := net.Listen("tcp", ":8099")
	if err != nil {
		panic(err.Error())
	}

	err = server.Serve(lis)
	if err != nil {
		panic(err.Error())
	}
}