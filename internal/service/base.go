package service

import (
	"github.com/ehwjh2010/cobra-example/internal/dao"
	"github.com/ehwjh2010/cobra-example/internal/proxy"
)

var OnStart = []func() error{
	dao.LoadDB,
	proxy.LoadCache,
}

var OnShutDown = []func() error{
	dao.CloseDB,
	proxy.CloseCache,
}
