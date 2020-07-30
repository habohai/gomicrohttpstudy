package weblib

import (
	"gomicrohttpstudy/services"

	"github.com/gin-gonic/gin"
)

func NewGinRouter(prodService services.ProdService) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(
		InitMiddleware(prodService),
		ErrorMiddleware(),
	)
	v1Group := ginRouter.Group("/v1")
	{
		v1Group.Handle("POST", "/prods", GetProdsList)
		v1Group.Handle("GET", "/prods/:pid", GetProdDetail)
	}
	return ginRouter
}
