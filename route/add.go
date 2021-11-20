package route

import (
	"github.com/ehwjh2010/cobra/http/response"
	"github.com/ehwjh2010/cobra/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Bind(engine *gin.Engine) {

	api := engine.Group("/api")

	{
		api.GET("/test", func(context *gin.Context) {
			log.Error("Invoke test api success", zap.String("name", "test"))
			response.Success(context, map[string]bool{
				"ok": true,
			})
		})
	}
}