package weblib

import (
	"context"
	"gomicrohttpstudy/services"
	"log"

	"github.com/gin-gonic/gin"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetProdDetail(ginCtx *gin.Context) {
	var prodReq services.ProdsRequest
	PanicIfError(ginCtx.BindUri(&prodReq))
	log.Println(prodReq.ProdId)
	prodService := ginCtx.Keys["prodservice"].(services.ProdService)
	resp, _ := prodService.GetProdsDetail(context.Background(), &prodReq)
	ginCtx.JSON(200, gin.H{"data": resp.Data})
}

func GetProdsList(ginCtx *gin.Context) {
	prodService := ginCtx.Keys["prodservice"].(services.ProdService)
	var prodReq services.ProdsRequest
	err := ginCtx.Bind(&prodReq)
	if err != nil {
		ginCtx.JSON(500, gin.H{"status": err.Error()})
		ginCtx.Abort()
	} else {
		prodRes, _ := prodService.GetProdsList(context.Background(), &prodReq)
		ginCtx.JSON(200, gin.H{"data": prodRes.Data})
	}
}
