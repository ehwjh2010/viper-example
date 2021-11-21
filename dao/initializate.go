package dao

import (
	"github.com/ehwjh2010/cobra/client"
	"github.com/ehwjh2010/cobra/db/cache"
	"github.com/ehwjh2010/cobra/db/rdb"
	"github.com/ehwjh2010/cobra/db/rdb/mysql"
	"log"
)

var (
	DBClient    *rdb.DBClient
	CacheClient *cache.RedisClient
)

//LoadDB 加载DB
func LoadDB(config *client.DB) {

	dbClient, err := mysql.InitMysql(config)

	if err != nil {
		log.Panicf("Load mysql failed!, err: %v", err)
	}

	DBClient = dbClient
}

//CloseDB 关闭DB
func CloseDB() error {
	return DBClient.Close()
}

//LoadCache 加载缓存
func LoadCache(config *client.Cache) {

	cacheClient, err := cache.InitCache(config)
	if err != nil {
		log.Panicf("Load redis failed!, err: %v\n", err)
	}

	CacheClient = cacheClient
}

//CloseCache 关闭缓存
func CloseCache() error {
	return CacheClient.Close()
}
