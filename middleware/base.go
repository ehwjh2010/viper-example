package middleware

import (
	"github.com/ehwjh2010/cobra/config"
	"github.com/ehwjh2010/cobra/extend/ginext/middleware"
	"github.com/gin-gonic/gin"
)

var Middlewares = []gin.HandlerFunc{
	middleware.AccessLog(nil, false, config.DefaultTimePattern),
	middleware.RecoveryWithZap(),
	middleware.Cors(middleware.OriginOpt("*")),
}
