package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/ehwjh2010/viper/client"
	"github.com/ehwjh2010/viper/helper/file"
	"github.com/ehwjh2010/viper/helper/path"
	"github.com/ehwjh2010/viper/helper/str"
	"github.com/ehwjh2010/viper/log"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Env             string       `yaml:"env" json:"env"`                         //环境标识
	Host            string       `yaml:"host" json:"host"`                       //地址
	Port            int          `yaml:"port" json:"port"`                       //端口
	ShutDownTimeout int          `yaml:"shutDownTimeout" json:"shutDownTimeout"` //优雅重启, 接收到相关信号后, 处理请求的最长时间, 单位: 秒， 默认5s
	Application     string       `yaml:"application" json:"application"`         //应用名
	Debug           bool         `yaml:"debug" json:"debug"`                     //debug
	Language        string       `yaml:"language" json:"language"`               //校验错误返回的语言
	Swagger         bool         `yaml:"swagger" json:"swagger"`                 //是否启动swagger
	EnableRtePool   bool         `yaml:"enableRtePool" json:"enableRtePool"`     //启用协程池
	LogConfig       client.Log   `yaml:"log" json:"log"`
	DBConfig        client.DB    `yaml:"db" json:"db"`
	CacheConfig     client.Cache `yaml:"cache" json:"cache"`
}

var Conf Config

//LoadConfig 从配置文件中加载配置
func init() {

	log.Debug("Start load config")

	configFilePath, err := ensureConfigPath()

	if err != nil {
		panic(err)
	}

	yamlFile, err := file.ReadFile(configFilePath)
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
	localConfigPath := path.PathJoin(currentDir, "config", "config_local.yaml")

	exist, err := path.EnsurePathExist(localConfigPath)

	if err != nil {
		return "", err
	}

	if exist {
		return localConfigPath, nil
	}

	//未读取到local配置文件, 则读取相应环境配置文件
	env := getEnv()

	configFileName := fmt.Sprintf("config_%s.yaml", strings.ToLower(env))

	configFilePath := path.PathJoin(currentDir, "config", configFileName)

	exist, err = path.EnsurePathExist(configFilePath)

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

	if str.IsEmpty(env) {
		env = "dev"
	}

	return env
}
