package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc_demo/grpc_stream_server_demo/proto"
	"net"
	"time"
)

type OrderMessage struct {}

func (om *OrderMessage) GetOrderInfo(req *proto.OrderRequest, orderInfo proto.OrderService_GetOrderInfoServer) error  {
	orderMap := map[string]proto.OrderInfo{
		"20211011": proto.OrderInfo{OrderId: "20211011",OrderName: "test01",OrderStatus: "successs"},
		"20211012": proto.OrderInfo{OrderId: "20211012",OrderName: "test02",OrderStatus: "successs"},
		"20211013": proto.OrderInfo{OrderId: "20211013",OrderName: "test03",OrderStatus: "successs"},
	}

	for id, ordInfo := range orderMap {
		if (req.TimeTmp <= time.Now().Unix()) {
			fmt.Println("订单序列号",id)
			fmt.Println("订单信息",ordInfo.OrderName,ordInfo.OrderStatus,ordInfo.OrderId)
			err := orderInfo.Send(&ordInfo)
			if err != nil {
				panic(err.Error())
			}
		}
	}
	return nil
}

func main() {

	orderMessage := new(OrderMessage)
	server := grpc.NewServer()
	proto.RegisterOrderServiceServer(server,orderMessage)

	lis, err := net.Listen("tcp", ":8099")
	if err != nil {
		panic(err.Error())
	}

	err = server.Serve(lis)
	if err != nil {
		panic(err.Error())
	}
}