package main

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro"
	etcd "github.com/micro/go-micro/registry/etcd"
	proto "github.com/rickynyairo/micro-framework-greeter/proto/greeter"
)

func main() {
	// Create a new service
	service := micro.NewService(
		micro.Name("greeter.client"),
		micro.Registry(etcd.NewRegistry()),
	)
	// Initialise the client and parse command line flags
	service.Init()

	// Create new greeter client
	greeter := proto.NewGreeterService("greeter", service.Client())

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.Request{Name: "Ricky"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Greeting)
}
