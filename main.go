package main

import (
	"github.com/ehwjh2010/cobra"
	"github.com/ehwjh2010/cobra-example/conf"
	"github.com/ehwjh2010/cobra-example/route"
	"github.com/ehwjh2010/cobra/enum"
	"github.com/ehwjh2010/cobra/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	middlewares := []gin.HandlerFunc{
		middleware.CobraZap(nil, true, enum.DefaultTimePattern),
		middleware.RecoveryWithZap(),
	}

	engine := cobra.Launch(conf.Conf.Application, conf.Conf.Mode, conf.Conf.LogConfig, middlewares)

	route.Bind(engine)

	cobra.Run(engine, conf.Conf.ServerConfig, nil, nil)
}

func init() {
	conf.LoadConfig()
}
