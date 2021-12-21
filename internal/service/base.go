package service

import (
	"github.com/ehwjh2010/viper-example/internal/dao"
	"github.com/ehwjh2010/viper-example/internal/proxy"
)

var OnStart = []func() error{
	dao.LoadDB,
	proxy.LoadCache,
}

var OnShutDown = []func() error{
	dao.CloseDB,
	proxy.CloseCache,
}
