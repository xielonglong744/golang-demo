package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_demo/grpcf_demo_three/proto"
	"time"
)

func main() {
	conn, err := grpc.Dial(":8099",grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()

	req := proto.OrderRequest{TimeTmp: time.Now().Unix(),OrderId: "20211011"}

	ordeClient := proto.NewOrderServiceClient(conn)

	orderInfo, err := ordeClient.GetOrderInfo(context.Background(),&req)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(orderInfo.OrderId,orderInfo.OrderStatus,orderInfo.OrderName)
}
