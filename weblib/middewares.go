package weblib

import (
	"gomicrohttpstudy/services"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(prodService services.ProdService) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Keys = make(map[string]interface{})
		context.Keys["prodservice"] = prodService
		context.Next()
	}
}
