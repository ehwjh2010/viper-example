package main

import (
	"github.com/ehwjh2010/viper-example/config"
	_ "github.com/ehwjh2010/viper-example/docs"
	"github.com/ehwjh2010/viper-example/internal/middleware"
	"github.com/ehwjh2010/viper-example/internal/service"
	"github.com/ehwjh2010/viper-example/router/demo"
	"github.com/ehwjh2010/viper/client"
	"github.com/ehwjh2010/viper/frame/ginext"
	"github.com/ehwjh2010/viper/util/object"
)

var setting client.Setting

// @title CobraExample API
// @version 1.0
// @description  Cobra使用示例
// @termsOfService http://swagger.io/terms/

// @contact.name Swagger API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:9090
// @BasePath /api
func main() {
	app := ginext.Viper(setting)

	demo.Bind(app.Engine())

	app.Run()
}

func init() {
	object.CopyProperties(&config.Conf, &setting)
	setting.Middlewares = middleware.Middlewares
	setting.OnStartUp = service.OnStart
	setting.OnShutDown = service.OnShutDown
}
