package weblib

import (
	"context"
	"gomicrohttpstudy/services"
	"strconv"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
)

func newProd(pid int32, pname string) *services.ProdModel {
	return &services.ProdModel{ProdID: pid, ProdName: pname}
}

func defaultProds() (*services.ProdListResponse, error) {
	models := make([]*services.ProdModel, 0)
	var i int32
	for i = 0; i < 5; i++ {
		models = append(models, newProd(200+i, strconv.Itoa(200+int(i))))
	}
	res := new(services.ProdListResponse)
	res.Data = models

	return res, nil
}

func GetProdsList(ginCtx *gin.Context) {
	prodService := ginCtx.Keys["prodservice"].(services.ProdService)
	var prodReq services.ProdsRequest
	err := ginCtx.Bind(&prodReq)
	if err != nil {
		ginCtx.JSON(500, gin.H{"status": err.Error()})
		ginCtx.Abort()
	} else {
		// 熔断代码改造
		// step1: 配置config
		configA := hystrix.CommandConfig{
			Timeout: 1000,
		}
		// step2: 配置command
		hystrix.ConfigureCommand("getprods", configA)

		var prodRes *services.ProdListResponse
		// step3: 执行，使用DO方法
		err := hystrix.Do(
			"getprods",
			func() error {
				prodRes, err = prodService.GetProdsList(context.Background(), &prodReq)
				return err
			},
			func(error) error {
				prodRes, err = defaultProds()
				return err
			},
		)

		if err != nil {
			ginCtx.JSON(500, gin.H{"status": err.Error()})
		} else {
			ginCtx.JSON(200, gin.H{"data": prodRes.Data})
		}
	}
}
