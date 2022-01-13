package demo

import (
	"github.com/ehwjh2010/viper-example/server/controller"
	"github.com/gin-gonic/gin"
)

func RegisterDemo(r *gin.RouterGroup) {

	{
		r.GET("/helloworld", controller.Helloworld)
		r.POST("/validate", controller.ValidateUser)
		r.GET("/task/count", controller.RoutineInfo)
		r.POST("/task", controller.BackgroundTask)
		r.GET("/request", controller.APIClient)
	}

	test := r.Group("/test")

	{
		test.GET("", controller.GetProjectConfig)
		test.GET("/:id", controller.QueryById)
		test.GET("/ids", controller.QueryByIds)
		test.GET("/cond", controller.QueryByCond)
		test.GET("/count", controller.QueryCountByCond)
		test.GET("/add", controller.AddRecord)
		test.GET("/update", controller.UpdateRecord)
	}

	cache := test.Group("/cache")

	{
		cache.GET("/set", controller.SetCache)
		cache.GET("/get", controller.GetCache)
	}
}
