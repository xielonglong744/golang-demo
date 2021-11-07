package main

import (
	"fmt"
	"grpc_demo/grpc_two/proto"
	"net/rpc"
	"time"
)

func main() {
	cli, err := rpc.DialHTTP("tcp",":9088")
	if err != nil {
		panic(err.Error())
	}
	timeTmp := time.Now().Unix()
	req := proto.OrderRequest{OrderId: "20211020",TimeTmp: timeTmp}
	var resp *proto.OrderInfo
	err = cli.Call("OrderService.GetOrderInfo",req,&resp)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(resp.OrderId,resp.OrderStatus,resp.OrderName)
}
