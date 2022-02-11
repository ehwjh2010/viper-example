package middleware

import (
	"github.com/ehwjh2010/viper/client/enums"
	"github.com/ehwjh2010/viper/frame/ginext/middleware"
	"github.com/gin-gonic/gin"
)

var Middlewares = []gin.HandlerFunc{
	middleware.AccessLog(nil, false, enums.DefaultTimePattern),
	middleware.RecoveryWithZap(),
	middleware.Cors(),
}
