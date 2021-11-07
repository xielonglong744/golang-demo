package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc_demo/grpc_steam_server/proto"
	"net"
	"time"
)

type OrderMessage struct {}

func (om *OrderMessage) GetOrderInfo(res *proto.OrderRequest,stream proto.OrderService_GetOrderInfoServer) error {
	orderMap := map[string]proto.OrderInfo{
		"20211011": proto.OrderInfo{OrderName: "test01",OrderStatus: "success",OrderId: "20211011"},
		"20211012": proto.OrderInfo{OrderName: "test02",OrderStatus: "success",OrderId: "20211012"},
		"20211013": proto.OrderInfo{OrderName: "test03",OrderStatus: "success",OrderId: "20211013"},
	}

	for id, info := range orderMap{
		if (res.TimeTmp <= time.Now().Unix()) {
			fmt.Println("订单序列号",id)
			fmt.Println("订单详情",info)

			stream.Send(&info)
		}
	}
	return nil
}

func main() {

	orderMessage := new(OrderMessage)
	server := grpc.NewServer()
	proto.RegisterOrderServiceServer(server,orderMessage)

	cli, err:= net.Listen("tcp", ":8099")
	if err != nil {
		panic(err.Error())
	}

	err = server.Serve(cli)
	if err != nil {
		panic(err.Error())
	}
}
