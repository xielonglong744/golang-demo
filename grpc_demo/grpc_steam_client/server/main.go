package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc_demo/grpc_steam_client/proto"
	"io"
	"net"
)

type OrderMessage struct {}

func (om *OrderMessage) AddOrderList(stream proto.ServiceOrder_AddOrderListServer) error  {
	fmt.Println("客户端流模式")
	for {
		orderRequest, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("结束接受")
			res := proto.OrderInfo{OrderId: "111111",OrderName: "1111111111",OrderStatus: "读取结束.............."}
			return  stream.SendAndClose(&res)
		}

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(orderRequest.OrderId,orderRequest.TimeTmp)
	}
}

func main() {
	orderMessage := new(OrderMessage)

	server := grpc.NewServer()
	proto.RegisterServiceOrderServer(server, orderMessage)

	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err.Error())
	}

	err = server.Serve(lis)
	if err != nil {
		panic(err.Error())
	}
}
