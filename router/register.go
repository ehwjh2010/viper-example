package router

import (
	"github.com/ehwjh2010/viper-example/config"
	v1 "github.com/ehwjh2010/viper-example/router/v1"
	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {

	api := engine.Group("/" + config.Conf.Application + "/api")

	v1.Bind(api)

}
