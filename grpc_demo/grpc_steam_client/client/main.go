package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_demo/grpc_steam_client/proto"
	"io"
	"time"
)

func main() {
	conn, err := grpc.Dial(":8090", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()

	orderMap := map[string]proto.OrderRequest{
		"20211011": proto.OrderRequest{OrderId: "20211011",TimeTmp: time.Now().Unix()},
		"20211012": proto.OrderRequest{OrderId: "20211012",TimeTmp: time.Now().Unix()},
		"20211013": proto.OrderRequest{OrderId: "20211013",TimeTmp: time.Now().Unix()},
	}
	serviceOrderClient := proto.NewServiceOrderClient(conn)
	listClient,err := serviceOrderClient.AddOrderList(context.Background())
	if err != nil {
		panic(err.Error())
	}
	for _, orderReq := range orderMap {
		err = listClient.Send(&orderReq)
		if err != nil {
			panic(err.Error())
		}
	}

	for {
		orderInfo, err := listClient.CloseAndRecv()
		if err == io.EOF {
			fmt.Println("读取结束")
			return
		}

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(orderInfo.OrderId,orderInfo.OrderName,orderInfo.OrderStatus)
	}
}
