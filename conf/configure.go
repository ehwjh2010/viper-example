package conf

import (
	"fmt"
	"github.com/ehwjh2010/cobra/client"
	"github.com/ehwjh2010/cobra/log"
	"github.com/ehwjh2010/cobra/util/fileutils"
	"github.com/ehwjh2010/cobra/util/pathutils"
	"github.com/ehwjh2010/cobra/util/strutils"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

type Config struct {
	Env             string       `yaml:"env" json:"env"`                         //环境标识
	Host            string       `yaml:"host" json:"host"`                       //地址
	Port            int          `yaml:"port" json:"port"`                       //端口
	ShutDownTimeout int          `yaml:"shutDownTimeout" json:"shutDownTimeout"` //优雅重启, 接收到相关信号后, 处理请求的最长时间, 单位: 秒， 默认5s
	Application     string       `yaml:"application" json:"application"`         //应用名
	Debug           bool         `yaml:"debug" json:"debug"`                     //debug
	Swagger         bool         `yaml:"swagger" json:"swagger"`                 //是否启动swagger
	LogConfig       client.Log   `yaml:"log" json:"log"`
	DBConfig        client.DB    `yaml:"db" json:"db"`
	CacheConfig     client.Cache `yaml:"cache" json:"cache"`
}

var Conf Config

//LoadConfig 从配置文件中加载配置
func LoadConfig() {

	log.Debug("Start load config")

	configFilePath, err := ensureConfigPath()

	if err != nil {
		panic(err)
	}

	yamlFile, err := fileutils.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		panic(err)
	}

	//设置环境标识
	env := getEnv()

	Conf.Env = env

	log.Debug("Load config success")
}

//ensureConfigPath 确定配置文件
func ensureConfigPath() (string, error) {
	currentDir, _ := os.Getwd()

	//优先读取本地配置, 利于本地开发以及线上配置
	localConfigPath := pathutils.PathJoin(currentDir, "conf", "config_local.yaml")

	exist, err := pathutils.EnsurePathExist(localConfigPath)

	if err != nil {
		return "", err
	}

	if exist {
		return localConfigPath, nil
	}

	//未读取到local配置文件, 则读取相应环境配置文件
	env := getEnv()

	configFileName := fmt.Sprintf("config_%s.yaml", strings.ToLower(env))

	configFilePath := pathutils.PathJoin(currentDir, "conf", configFileName)

	exist, err = pathutils.EnsurePathExist(configFilePath)

	if err != nil {
		return "", err
	} else if !exist {
		return "", err
	}

	return configFilePath, nil
}

//getEnv 获取环境标识
func getEnv() string {
	env := strings.ToLower(os.Getenv("ENV"))

	if strutils.IsEmptyStr(env) {
		env = "dev"
	}

	return env
}
