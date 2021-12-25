package router

import (
	"github.com/ehwjh2010/viper-example/router/demo"
	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {

	demo.RegisterDemo(engine)
}
