package middleware

import (
	"github.com/ehwjh2010/cobra/config"
	"github.com/ehwjh2010/cobra/middleware"
	"github.com/gin-gonic/gin"
)

var Middllewares = []gin.HandlerFunc{
	middleware.CobraZap(nil, false, config.DefaultTimePattern),
	middleware.RecoveryWithZap(),
}
