package v1

import (
	"github.com/ehwjh2010/viper-example/router/v1/demo"
	"github.com/gin-gonic/gin"
)

func Bind(api *gin.RouterGroup) {

	v1 := api.Group("/v1")
	demo.RegisterDemo(v1)
}
