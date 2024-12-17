package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gvb-server/config"
	"log"
	"os"
)

// InitConf 读取yaml配置文件
func InitConf() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error:%s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("unmarshal yamlConf error:%s", err)
	}
	log.Println("config yamlFile load Init success.")
	fmt.Println(c)
}
