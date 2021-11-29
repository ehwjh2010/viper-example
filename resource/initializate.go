package resource

import (
	"github.com/ehwjh2010/cobra-example/conf"
	"github.com/ehwjh2010/cobra/db/cache"
	"github.com/ehwjh2010/cobra/db/rdb"
	"github.com/ehwjh2010/cobra/db/rdb/mysql"
)

var (
	DBClient    *rdb.DBClient
	CacheClient *cache.RedisClient
)

//LoadDB 加载DB
func LoadDB() error {

	dbClient, err := mysql.InitMysql(&conf.Conf.DBConfig)

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
func LoadCache() error {

	cacheClient, err := cache.InitCache(&conf.Conf.CacheConfig)
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
