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
		Timeout: 1000,
	}

	hystrix.ConfigureCommand(cmdName, configA)

	return hystrix.Do(
		cmdName,
		func() error {
			return l.Client.Call(ctx, req, rsp)
		},
		func(err error) error {
			defaultProds(rsp)
			return nil
		},
	)
}

func NewProdsWrapper(c client.Client) client.Client {
	return &ProdsWrapper{c}
}
