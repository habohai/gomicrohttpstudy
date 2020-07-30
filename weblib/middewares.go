package weblib

import (
	"fmt"
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

func ErrorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				context.JSON(500, gin.H{"status": fmt.Sprintf("%s", r)})
				context.Abort()
			}
		}()
		context.Next()
	}
}
