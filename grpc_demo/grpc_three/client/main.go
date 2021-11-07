package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_demo/grpc_three/proto"
	"time"
)

func main() {

	conn, err := grpc.Dial(":8099",grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()

	ordClient :=proto.NewOrderServiceClient(conn)

	req := &proto.OrderRequest{OrderId: "20211011",TimeTmp: time.Now().Unix()}

	orderInfo, err := ordClient.GetOrderInfo(context.Background(),req)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(orderInfo.OrderId,orderInfo.OrderName,orderInfo.OrderStatus)
}