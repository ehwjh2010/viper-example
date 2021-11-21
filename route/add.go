package route

import (
	"github.com/ehwjh2010/cobra/http/response"
	"github.com/gin-gonic/gin"
)

func Bind(engine *gin.Engine) {

	api := engine.Group("/api")

	{
		api.GET("/test", func(context *gin.Context) {
			response.Success(context, map[string]bool{
				"ok": true,
			})
		})
	}
}