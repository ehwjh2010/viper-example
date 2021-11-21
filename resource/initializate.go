package resource

import (
	"github.com/ehwjh2010/cobra/client"
	"github.com/ehwjh2010/cobra/db/cache"
	"github.com/ehwjh2010/cobra/db/rdb"
	"github.com/ehwjh2010/cobra/db/rdb/mysql"
)

var (
	DBClient    *rdb.DBClient
	CacheClient *cache.RedisClient
)

//LoadDB 加载DB
func LoadDB(config *client.DB) error {

	dbClient, err := mysql.InitMysql(config)

	if err != nil {
		return err
	}

	DBClient = dbClient
	return nil
}

//CloseDB 关闭DB
func CloseDB() error {
	return DBClient.Close()
}

//LoadCache 加载缓存
func LoadCache(config *client.Cache) error {

	cacheClient, err := cache.InitCache(config)
	if err != nil {
		return err
	}

	CacheClient = cacheClient
	return nil
}

//CloseCache 关闭缓存
func CloseCache() error {
	return CacheClient.Close()
}
