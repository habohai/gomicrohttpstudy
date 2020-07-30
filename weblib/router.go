package weblib

import (
	"gomicrohttpstudy/services"

	"github.com/gin-gonic/gin"
)

func NewGinRouter(prodService services.ProdService) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(InitMiddleware(prodService))
	v1Group := ginRouter.Group("/v1")
	{
		v1Group.Handle("POST", "/prods", GetProdsList)
	}
	return ginRouter
}
