package proxy

import (
	"github.com/ehwjh2010/cobra-example/configs"
	"github.com/ehwjh2010/viper/db/cache"
)

var RedisClient *cache.RedisClient

//LoadCache 载入Redis
func LoadCache() error {
	client, err := cache.InitCache(&configs.Conf.CacheConfig)
	if err != nil {
		return err
	}

	RedisClient = client
	return nil
}

//CloseCache 关闭Redis
func CloseCache() error {
	return RedisClient.Close()
}
