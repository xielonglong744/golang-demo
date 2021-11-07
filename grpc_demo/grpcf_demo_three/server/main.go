package main

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"grpc_demo/grpcf_demo_three/proto"
	"net"
	"time"
)

type OrderMessage struct{}

func (om *OrderMessage) GetOrderInfo(ctx context.Context, request *proto.OrderRequest) (*proto.OrderInfo, error) {
	orderMap := map[string]proto.OrderInfo{
		"20211011": proto.OrderInfo{OrderId: "20211011", OrderName: "test01", OrderStatus: "success"},
		"20211012": proto.OrderInfo{OrderId: "20211012", OrderName: "test02", OrderStatus: "failured"},
		"20211013": proto.OrderInfo{OrderId: "20211013", OrderName: "test03", OrderStatus: "success"},
	}

	current := time.Now().Unix()

	var resp *proto.OrderInfo

	if request.TimeTmp > current {
		*resp = proto.OrderInfo{OrderId: "", OrderName: "", OrderStatus: "订单异常"}
	} else {
		res := orderMap[request.OrderId]
		if res.OrderId != "" {
			return &res, nil
		} else {
			return nil, errors.New("server error")
		}
	}
	return resp, nil
}

func main() {
	orderMessage := new(OrderMessage)
	server := grpc.NewServer()
	proto.RegisterOrderServiceServer(server, orderMessage)

	lis, err := net.Listen("tcp", ":8099")
	if err != nil {
		panic(err.Error())
	}
	err = server.Serve(lis)
	if err != nil {
		panic(err.Error())
	}
}
