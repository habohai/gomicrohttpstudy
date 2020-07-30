package main

import (
	"context"
	"fmt"
	"gomicrohttpstudy/services"
	"gomicrohttpstudy/weblib"
	"gomicrohttpstudy/wrapper"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/web"
)

type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Endpoint())
	return l.Client.Call(ctx, req, rsp)
}

func NewLogWrapper(c client.Client) client.Client {
	return &logWrapper{c}
}

func main() {
	reg := etcd.NewRegistry(
		registry.Addrs("192.168.31.82:12379"),
	)

	myService := micro.NewService(
		micro.Name("prodsservic.client"),
		micro.WrapClient(NewLogWrapper),
		micro.WrapClient(wrapper.NewProdsWrapper),
	)

	prodService := services.NewProdService("prodservice", myService.Client())

	httpServer := web.NewService(
		web.Name("httpprodservice"),
		web.Address(":9061"),
		web.Handler(weblib.NewGinRouter(prodService)),
		web.Registry(reg),
		web.Metadata(map[string]string{"protocol": "http"}),
	)

	httpServer.Init()
	httpServer.Run()
}
