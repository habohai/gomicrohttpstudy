package weblib

import (
	"context"
	"gomicrohttpstudy/services"

	"github.com/gin-gonic/gin"
)

func GetProdsList(ginCtx *gin.Context) {
	prodService := ginCtx.Keys["prodservice"].(services.ProdService)
	var prodReq services.ProdsRequest
	err := ginCtx.Bind(&prodReq)
	if err != nil {
		ginCtx.JSON(500, gin.H{"status": err.Error()})
		ginCtx.Abort()
	} else {
		prodRes, err := prodService.GetProdsList(context.Background(), &prodReq)
		if err != nil {
			ginCtx.JSON(500, gin.H{"status": err.Error()})
		} else {
			ginCtx.JSON(200, gin.H{"data": prodRes.Data})
		}
	}
}
