package main

import (
	"github.com/ehwjh2010/cobra"
	"github.com/ehwjh2010/cobra-example/conf"
	"github.com/ehwjh2010/cobra-example/middleware"
	"github.com/ehwjh2010/cobra-example/service"
	"github.com/ehwjh2010/cobra/client"
	"github.com/ehwjh2010/cobra/util/structutils"
)

var setting client.Setting

// @title CobraExample接口文档
// @version 1.0
// @description CobraExample接口文档

// @contact.name API Support
// @contact.url http://www.swagger.io/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api
func main() {
	app := cobra.Cobra(setting)

	Bind(app.Engine)

	app.Run()
}

func init() {
	conf.LoadConfig()
	structutils.CopyProperties(&conf.Conf, &setting)
	setting.Middlewares = middleware.Middlewares
	setting.OnStartUp = service.OnStart
	setting.OnShutDown = service.OnShutDown
}
