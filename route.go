package main

import (
	"github.com/ehwjh2010/cobra-example/controller"
	"github.com/gin-gonic/gin"
)

func Bind(engine *gin.Engine) {

	api := engine.Group("/api")

	{
		api.GET("/helloworld", controller.Helloworld)
	}
	
	test := api.Group("/test")
	
	{
		test.GET("", controller.GetProjectConfig)
		test.GET("/:id", controller.QueryById)
		test.GET("/ids", controller.QueryByIds)
		test.GET("/cond", controller.QueryByCond)
		test.GET("/add", controller.AddRecord)
	}
}
