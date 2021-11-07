package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc_demo/grpc_stream_client_demo/proto"
	"io"
	"net"
)

type OrderMessage struct {}



func (om *OrderMessage) AddOrderList(stream proto.OrderService_AddOrderListServer) error  {
	fmt.Println("开启客户端流模式")

	for {
		orderRequest, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("数据接受结束")
			res := proto.OrderInfo{OrderStatus: "订单结束", OrderId: "111111111",OrderName: "xielonglong"}
			return stream.SendAndClose(&res)

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
	proto.RegisterOrderServiceServer(server,orderMessage)

	lis ,err := net.Listen("tcp",":8099")
	if err != nil {
		panic(err.Error())
	}
	err = server.Serve(lis)
	if err != nil {
		panic(err.Error())
	}
}
