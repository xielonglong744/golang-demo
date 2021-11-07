package main

import (
	"fmt"
	"net/rpc"
)

//import (
//	"fmt"
//	"net/rpc"
//)
//
//
//func main() {
//
//	client, err := rpc.DialHTTP("tcp", "localhost:8081")
//	if err != nil {
//		panic(err.Error())
//	}
//
//	var req float32 //请求值
//	req = 3
//
//	var resp *float32 //返回值
//	resp = new(float32)
//	err = client.Call("MathUtil.CircularArea", req, resp)
//	if err != nil {
//		panic(err.Error())
//	}
//	fmt.Println(*resp)
//}

func main() {
	client, err := rpc.DialHTTP("tcp","localhost:9099")
	if err != nil {
		panic(err.Error())
	}
	var req float32 = 3
	//var resp = new(float32)

	// 同步调用
	//err = client.Call("Mutal.CirAeal",req,resp)
	//if err != nil {
	//	panic(err.Error())
	//}
	//fmt.Println(*resp)

	// 异步调用
	//var syncRep = new(*float32)
	var syncRep *float32
	cli := client.Go("Mutal.CirAeal",req,&syncRep,nil)
	replDone := <- cli.Done
	fmt.Println(replDone)
	fmt.Println(*syncRep)
}





























