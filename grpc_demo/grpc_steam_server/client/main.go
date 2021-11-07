package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_demo/grpc_steam_server/proto"
	"io"
	"time"
)

func main() {
	conn, err  := grpc.Dial(":8099",grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()

	orderServiceClient := proto.NewOrderServiceClient(conn)
	req := proto.OrderRequest{OrderId: "20211011",TimeTmp: time.Now().Unix()}
	getInfoClient, err := orderServiceClient.GetOrderInfo(context.Background(),&req)
	if err != nil {
		panic(err.Error())
	}
	for {
		orderInfo, err := getInfoClient.Recv()
		if err == io.EOF {
			fmt.Println("读取结束")
			return
		}

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(orderInfo.OrderName,orderInfo.OrderId,orderInfo.OrderId)
	}

}
