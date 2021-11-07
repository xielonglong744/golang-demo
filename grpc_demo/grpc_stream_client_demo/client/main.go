package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_demo/grpc_stream_client_demo/proto"
	"io"
	"time"
)

func main() {
	conn, err := grpc.Dial(":8099", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()

	ordSerClient :=proto.NewOrderServiceClient(conn)
	orderMap := map[string]proto.OrderRequest{
		"20211011": proto.OrderRequest{TimeTmp: time.Now().Unix(),OrderId: "20211011"},
		"20211012": proto.OrderRequest{TimeTmp: time.Now().Unix(),OrderId: "20211012"},
		"20211013": proto.OrderRequest{TimeTmp: time.Now().Unix(),OrderId: "20211013"},
		"20211014": proto.OrderRequest{TimeTmp: time.Now().Unix(),OrderId: "20211014"},
	}


	ordSerAddClient, err := ordSerClient.AddOrderList(context.Background())
	if err != nil {
		panic(err.Error())
	}

	for _, orderReq := range orderMap{
		err = ordSerAddClient.Send(&orderReq)
		if err != nil {
			panic(err.Error())
		}

	}

	for {
		orderInfo, err := ordSerAddClient.CloseAndRecv()
		if err == io.EOF {
			fmt.Println("读取数据结束")
			return
		}
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(orderInfo.OrderId,orderInfo.OrderName,orderInfo.OrderStatus)
	}
}
