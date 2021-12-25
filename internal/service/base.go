package service

import (
	"github.com/ehwjh2010/viper-example/internal/dao"
	"github.com/ehwjh2010/viper-example/internal/proxy/cache"
)

var OnStart = []func() error{
	dao.LoadDB,
	cache.LoadCache,
}

var OnShutDown = []func() error{
	dao.CloseDB,
	cache.CloseCache,
}
