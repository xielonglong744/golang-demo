package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_demo/grpc_stream_server_demo/proto"
	"io"
	"time"
)

func main() {
	conn, err := grpc.Dial(":8099",grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()

	orderSercieClient := proto.NewOrderServiceClient(conn)

	var req = proto.OrderRequest{OrderId: "20211011",TimeTmp: time.Now().Unix()}
	infoClient,err := orderSercieClient.GetOrderInfo(context.Background(),&req)
	if err != nil {
		panic(err.Error())
	}

	ordInfo, err := infoClient.Recv()
	if err == io.EOF {
		fmt.Println("读取结束")
		return
	}

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(ordInfo.OrderName,ordInfo.OrderStatus,ordInfo.OrderId)
}
