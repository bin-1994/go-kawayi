package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type ConfModel struct {
	Base struct {
		Model string
		Port  string
		Proxy string
	}
	Db struct {
		Type     string
		Host     string
		Database string
		Username string
		Password string
		Port     int
		Chaset   string
	}
	Redis struct {
		Host     string
		Port     int
		Password string
	}
	Lsm struct {
		Verify string
		Key    string
	}
}

var YamlConf ConfModel

func InitConf() {
	// 读取配置文件
	conf, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		fmt.Println("ReadFile:", err.Error())
		os.Exit(-1)
	}
	if err := yaml.Unmarshal(conf, &YamlConf); err != nil {
		fmt.Println("invalid config: ", err.Error())
		os.Exit(-1)
	}
}
