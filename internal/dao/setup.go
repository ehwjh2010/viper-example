package dao

import (
	"github.com/ehwjh2010/viper-example/config"
	"github.com/ehwjh2010/viper/db/rdb"
	"github.com/ehwjh2010/viper/db/rdb/mysql"
)

var DBClient *rdb.DBClient

//LoadDB 加载DB
func LoadDB() error {

	dbClient, err := mysql.SetUp(config.Conf.DBConfig)

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
