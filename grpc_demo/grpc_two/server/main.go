package main

import (
	"errors"
	"grpc_demo/grpc_two/proto"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type OrderService struct {}

func (os *OrderService) GetOrderInfo(req proto.OrderRequest, resp *proto.OrderInfo) error  {
	orderMap := map[string]proto.OrderInfo{
		"20211020": proto.OrderInfo{OrderId: "20211020",OrderName: "test01",OrderStatus: "成功"},
		"20211021": proto.OrderInfo{OrderId: "20211021",OrderName: "test02",OrderStatus: "失败"},
		"20211022": proto.OrderInfo{OrderId: "20211021",OrderName: "test03",OrderStatus: "成功"},
	}

	current := time.Now().Unix()

	if req.TimeTmp > current {
		*resp = proto.OrderInfo{OrderId: "0",OrderName: "",OrderStatus: "订单异常"}
	}else {
		res := orderMap[req.OrderId]
		if res.OrderId != "" {
			*resp = orderMap[res.OrderId]
		}else {
			return errors.New("server error")
		}
	}

	return nil
}

func main() {
	orderService := new(OrderService)
	rpc.Register(orderService)
	rpc.HandleHTTP()

	lis, err := net.Listen("tcp",":9088")
	if err != nil {
		panic(err.Error())
	}
	http.Serve(lis,nil)
}