package middleware

import (
	"github.com/ehwjh2010/viper/frame/ginext/middleware"
	"github.com/ehwjh2010/viper/global"
	"github.com/gin-gonic/gin"
)

var Middlewares = []gin.HandlerFunc{
	middleware.AccessLog(nil, false, global.DefaultTimePattern),
	middleware.RecoveryWithZap(),
	middleware.Cors(),
}
