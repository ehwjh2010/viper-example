package cache

import (
	"github.com/ehwjh2010/cobra-example/configs"
	"github.com/ehwjh2010/cobra/db/cache"
)

var RedisClient *cache.RedisClient

//LoadCache 加载缓存
func LoadCache() error {

	cacheClient, err := cache.InitCache(&configs.Conf.CacheConfig)
	if err != nil {
		return err
	}

	RedisClient = cacheClient
	return nil
}

//CloseCache 关闭缓存
func CloseCache() error {
	return RedisClient.Close()
}
