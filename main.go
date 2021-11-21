package main

import (
	"github.com/ehwjh2010/cobra"
	"github.com/ehwjh2010/cobra-example/conf"
	"github.com/ehwjh2010/cobra-example/middleware"
	"github.com/ehwjh2010/cobra-example/service"
)

func main() {
	app := cobra.Cobra(conf.Conf.Application, conf.Conf.Debug, conf.Conf.LogConfig, middleware.Middlewares)

	Bind(app.Engine)

	app.Run(conf.Conf.ServerConfig, service.OnStart, service.OnShutDown)

}

func init() {
	conf.LoadConfig()
}
