package service

import (
	"github.com/ehwjh2010/cobra-example/resource"
)


var OnStart = []func() error {
	resource.LoadDB,
	resource.LoadCache,
}

var OnShutDown = []func() error {
	resource.CloseDB,
	resource.CloseCache,
}