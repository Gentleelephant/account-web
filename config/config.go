package config

import (
	"github.com/Gentleelephant/common/consts"
	"github.com/Gentleelephant/common/utils"
	"github.com/spf13/viper"
)

var (
	LocalConfig *viper.Viper
	FilePath    string // 配置文件路径，命令启动时指定
	NacosConfig *utils.NacosConfigparams
)

func GetLocalConfig() {
	localConfig, err := utils.GetConfig(FilePath, consts.ConfigFileType)
	if err != nil {
		panic(err)
	}
	LocalConfig = localConfig
}
