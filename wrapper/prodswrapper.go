package wrapper

import (
	"context"
	"gomicrohttpstudy/services"
	"strconv"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"
)

func newProd(pid int32, pname string) *services.ProdModel {
	return &services.ProdModel{ProdID: pid, ProdName: pname}
}

// 通用商品降级方法
func defaultData(rsp interface{}) {
	switch t := rsp.(type) {
	case *services.ProdListResponse:
		defaultProds(rsp)
	case *services.ProdDetailResponse:
		t.Data = newProd(10, "降级商品")
	}
}

// 商品列表降级方法
func defaultProds(rsp interface{}) {
	models := make([]*services.ProdModel, 0)
	var i int32
	for i = 0; i < 5; i++ {
		models = append(models, newProd(200+i, strconv.Itoa(200+int(i))))
	}
	result := rsp.(*services.ProdListResponse)
	result.Data = models
}

type ProdsWrapper struct {
	client.Client
}

func (l *ProdsWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint()

	configA := hystrix.CommandConfig{
		Timeout:                3000,
		RequestVolumeThreshold: 2,
		ErrorPercentThreshold:  50,
		SleepWindow:            5000,
	}

	hystrix.ConfigureCommand(cmdName, configA)

	return hystrix.Do(
		cmdName,
		func() error {
			return l.Client.Call(ctx, req, rsp)
		},
		func(err error) error {
			defaultData(rsp)
			return nil
		},
	)
}

func NewProdsWrapper(c client.Client) client.Client {
	return &ProdsWrapper{c}
}
