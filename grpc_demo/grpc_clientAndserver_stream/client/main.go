package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_demo/grpc_clientAndserver_stream/proto"
	"io"
	"time"
)

func main() {
	conn, err := grpc.Dial(":8099", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()

	orderServiceClient := proto.NewOrderServiceClient(conn)
	orderSerClient, err := orderServiceClient.GetOrderInfo(context.Background())
	if err != nil {
		panic(err.Error())
	}

	orderReq := map[string]proto.OrderRequest{
		"20211011": proto.OrderRequest{OrderId: "20211011",TimeTmp: time.Now().Unix()},
		"20211012": proto.OrderRequest{OrderId: "20211012",TimeTmp: time.Now().Unix()},
		"20211013": proto.OrderRequest{OrderId: "20211013",TimeTmp: time.Now().Unix()},
		"20211014": proto.OrderRequest{OrderId: "20211014",TimeTmp: time.Now().Unix()},
	}

	for _,req := range orderReq {
		err = orderSerClient.Send(&req,)
		if err != nil {
			panic(err.Error())
		}
	}
	err = orderSerClient.CloseSend()
	if err != nil {
		panic(err.Error())
	}

	for {
		orderInfo, err := orderSerClient.Recv()
		if err == io.EOF {
			fmt.Println("读取数据结束")
			return
		}

		if err != nil {
			return
		}

		fmt.Println(orderInfo.OrderId,orderInfo.OrderName,orderInfo.OrderStatus)
	}

}
