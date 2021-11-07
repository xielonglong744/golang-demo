package main

import (
	"fmt"
	"grpc_demo/grpc_demo_two/proto"
	"net/rpc"
	"time"
)

func main() {
	timeTmp := time.Now().Unix()

	req := proto.OrderRequest{
		OrderId: "20211012",
		Timetmp: timeTmp,
	}

	var resp *proto.OrderInfo

	cli, err := rpc.DialHTTP("tcp",":8081")
	if err != nil {
		panic(err.Error())
	}

	err = cli.Call("OrderMessage.GetOrderInfo",req,&resp)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(resp.OrderName,resp.OrderStatus,resp.Orderid)
}
