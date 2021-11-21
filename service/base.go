package service

import (
	"github.com/ehwjh2010/cobra-example/conf"
	"github.com/ehwjh2010/cobra-example/resource"
	"github.com/ehwjh2010/cobra/types"
)

func OnStart() error {
	var multiErr types.MultiErr

	multiErr.AddErr(resource.LoadDB(&conf.Conf.DBConfig), resource.LoadCache(&conf.Conf.CacheConfig))

	if multiErr.IsEmpty() {
		return nil
	}

	return &multiErr
}

var OnShutDown = []func() error {
	resource.CloseDB,
	resource.CloseCache,
}