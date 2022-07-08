package config

import (
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

var FilePath string
var Config AccountConfig

func Load() error {
	fmt.Println("Config init:", Config)
	if FilePath == "" {
		return errors.New("config file path is not empty")
	}
	fmt.Println("load config file:", FilePath)
	v := viper.New()
	v.SetConfigFile(FilePath)
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	if err := v.Unmarshal(&Config); err != nil {
		return err
	}
	fmt.Println("Config:", Config)

	// 开启实时监控
	v.WatchConfig()

	// 文件更新的回调函数
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("监听到配置改变,重新加载配置")
		if err := v.Unmarshal(&Config); err != nil {
			return
		}
	})
	return nil
}

func LoadYaml() {
	open, err := os.Open(FilePath)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	all, err := ioutil.ReadAll(open)
	if err != nil {
		fmt.Println("read file err:", err)
		return
	}
	err = yaml.Unmarshal(all, &Config)
	if err != nil {
		fmt.Println("unmarshal err:", err)
		return
	}
}
