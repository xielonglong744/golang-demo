package main

import (
	"context"
	"github.com/micro/go-micro"
	pb "micro/helloworld/proto"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("hello"),
	)

	service.Init()

	err := pb.RegisterGreeterHandler(service.Server(), new(Greeter))
	if err != nil {
		panic(err.Error())
	}

	err = service.Run()
	if err != nil {
		panic(err.Error())
	}
}
