package main

import (
	"errors"
	"grpc_demo/grpc_demo_two/proto"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type OrderMessage struct {}

func (om *OrderMessage) GetOrderInfo(req proto.OrderRequest, resp *proto.OrderInfo) error {
	orderMap := map[string]proto.OrderInfo{
		"20211011": proto.OrderInfo{Orderid: "20211011",OrderName: "test",OrderStatus: "success"},
		"20211012": proto.OrderInfo{Orderid: "20211012",OrderName: "test2",OrderStatus: "failure"},
		"20211013": proto.OrderInfo{Orderid: "20211013",OrderName: "test3",OrderStatus: "unknown"},
	}

	current := time.Now().Unix()

	if (req.Timetmp > current) {
		*resp = proto.OrderInfo{Orderid: "",OrderName: "",OrderStatus: "订单无效"}
	} else {
		res := orderMap[req.OrderId]
		if res.Orderid != "" {
			*resp = orderMap[res.Orderid]
		} else {
			return errors.New("server error")
		}
	}
	return nil
}

func main() {
	orderMessage := new(OrderMessage)
	rpc.Register(orderMessage)

	rpc.HandleHTTP()

	cli, err := net.Listen("tcp",":8081")
	if err != nil {
		panic(err.Error())
	}

	http.Serve(cli, nil)
}