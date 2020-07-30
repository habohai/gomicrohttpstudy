package main

import (
	"gomicrohttpstudy/services"
	"gomicrohttpstudy/weblib"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/web"
)

func main() {
	reg := etcd.NewRegistry(
		registry.Addrs("192.168.31.82:12379"),
	)

	myService := micro.NewService(
		micro.Name("prodsservic.client"),
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
