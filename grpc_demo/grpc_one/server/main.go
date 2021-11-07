package main

import (
	"math"
	"net/http"
	"net/rpc"
)

//import (
//	"math"
//	"net/http"
//	"net/rpc"
//)
//
//// 定义服务及暴露服务
//type MathUtil struct {}
//
//// 定义方法
//func (m *MathUtil) CircularArea(req float32, resp *float32) error {
//	*resp = math.Pi * req * req
//	return nil
//}
//
//func main() {
//	// 初始化指针类型
//	mathUtil := new(MathUtil)
//
//	// 服务注册
//	err := rpc.Register(mathUtil)
//	if err != nil {
//		panic(err.Error())
//	}
//   // 将服务注册到http协议上 方便调用者可以使用http协议调用
//	rpc.HandleHTTP()
//
//	// 在特定的端口进行监听
//	//listter, err := net.Listen("tcp",":8081")
//	//if err != nil {
//	//	panic(err.Error())
//	//}
//	//http.Serve(listter,nil)
//	http.ListenAndServe(":8081",nil)
//}

type Mutal struct {}

func (m *Mutal) CirAeal(req float32, resp *float32) error  {
	*resp = math.Pi * req * req
	return nil
}

func main() {
	mutail := new(Mutal)
	err := rpc.Register(mutail)
	if err != nil {
		panic(err.Error())
	}

	rpc.HandleHTTP()

	http.ListenAndServe(":9099", nil)
}















