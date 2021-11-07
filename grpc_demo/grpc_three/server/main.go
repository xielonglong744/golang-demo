package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc_demo/grpc_three/proto"
	"net"
	"time"
)

type OrderMessage struct {}

func (om *OrderMessage) GetOrderInfo(ctx context.Context, req *proto.OrderRequest) (*proto.OrderInfo,error)  {
	orderMap := map[string]proto.OrderInfo{
		"20211011": proto.OrderInfo{OrderId: "20211011",OrderName: "test01",OrderStatus: "success"},
		"20211012": proto.OrderInfo{OrderId: "20211012",OrderName: "test02",OrderStatus: "success"},
		"20211013": proto.OrderInfo{OrderId: "20211013",OrderName: "test03",OrderStatus: "success"},
	}

	currrent := time.Now().Unix()
	var resp proto.OrderInfo
	if (req.TimeTmp > currrent) {
		resp = proto.OrderInfo{OrderId: "",OrderName: "",OrderStatus: "订单异常"}
	} else {
		res := orderMap[req.OrderId]
		if res.OrderId != "" {
			resp = orderMap[req.OrderId]
		}
	}
	return &resp,nil
}

func main() {
	ordermessage := new(OrderMessage)

	server := grpc.NewServer()
	proto.RegisterOrderServiceServer(server,ordermessage)

	lis,err := net.Listen("tcp",":8099")
	if err != nil {
		panic(err.Error())
	}

	server.Serve(lis)
}