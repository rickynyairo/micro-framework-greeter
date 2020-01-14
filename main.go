package main

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro"
	proto "github.com/rickynyairo/micro-framework-greeter/proto/greeter"
	etcd "github.com/micro/go-micro/registry/etcd"
)

// Greeter implements greeter service
type Greeter struct{}

// Hello function of Greeter service
func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Registry(etcd.NewRegistry()),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
