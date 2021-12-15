package service

import (
	"github.com/ehwjh2010/cobra-example/internal/cache"
	"github.com/ehwjh2010/cobra-example/internal/dao"
)

var OnStart = []func() error{
	dao.LoadDB,
	cache.LoadCache,
}

var OnShutDown = []func() error{
	dao.CloseDB,
	cache.CloseCache,
}
