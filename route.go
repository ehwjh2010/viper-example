package main

import (
	"github.com/ehwjh2010/cobra-example/controller"
	"github.com/gin-gonic/gin"
)

func Bind(engine *gin.Engine) {

	api := engine.Group("/api")

	{
		api.GET("/helloworld", controller.Helloworld)
		api.GET("/test", controller.GetProjectConfig)
		api.GET("/test/:id", controller.QueryById)
		api.GET("/test/ids", controller.QueryByIds)
		api.GET("/test/add", controller.AddRecord)
	}
}
