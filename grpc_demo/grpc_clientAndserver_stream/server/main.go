package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc_demo/grpc_clientAndserver_stream/proto"
	"io"
	"net"
	"time"
)

type OrderMessage struct{}

func (om *OrderMessage) GetOrderInfo(stream proto.OrderService_GetOrderInfoServer) error {
	for {
		orderReq, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("读取数据结束")
			return err
		}

		if err != nil {
			panic(err.Error())
		}
		fmt.Println(orderReq.OrderId,orderReq.TimeTmp)

		orderMap := map[string]proto.OrderInfo{
			"20211011": proto.OrderInfo{OrderId: "20211011", OrderStatus: "successs", OrderName: "test01"},
			"20211012": proto.OrderInfo{OrderId: "20211012", OrderStatus: "successs", OrderName: "test02"},
			"20211013": proto.OrderInfo{OrderId: "20211013", OrderStatus: "successs", OrderName: "test03"},
			"20211014": proto.OrderInfo{OrderId: "20211014", OrderStatus: "successs", OrderName: "test04"},
		}
		if orderReq.TimeTmp <= time.Now().Unix() {
			res := orderMap[orderReq.OrderId]
			if res.OrderId != "" {
				err = stream.Send(&res)
				if err == io.EOF {
					fmt.Println("发送数据结束")
					return err
				}
				if err != nil {
					panic(err.Error())
				}
			}
		}

	}
	return nil
}

func main() {

	orderMessage := new(OrderMessage)

	server := grpc.NewServer()

	proto.RegisterOrderServiceServer(server,orderMessage)

	lis, err := net.Listen("tcp",":8099")
	if err != nil {
		panic(err.Error())
	}

	err = server.Serve(lis)
	if err != nil {
		panic(err.Error())
	}
}
