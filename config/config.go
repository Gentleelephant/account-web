package config

import (
	"github.com/Gentleelephant/common/consts"
	"github.com/Gentleelephant/common/utils"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
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

	NacosConfig = &utils.NacosConfigparams{
		DataId: localConfig.GetString(consts.NacosDataId),
		Group:  localConfig.GetString(consts.NacosGroup),
		ClientConfig: constant.ClientConfig{
			TimeoutMs:    localConfig.GetUint64(consts.NacosTimeoutMs),
			BeatInterval: localConfig.GetInt64(consts.NacosBeatInterval),
			NamespaceId:  localConfig.GetString(consts.NacosNamespaceId),
			CacheDir:     localConfig.GetString(consts.NacosCacheDir),
			LogDir:       localConfig.GetString(consts.NacosLogDir),
			LogLevel:     localConfig.GetString(consts.NacosLogLevel),
			ContextPath:  localConfig.GetString(consts.NacosContextPath),
		},
		ServerConfig: constant.ServerConfig{
			Scheme:      "http",
			ContextPath: localConfig.GetString(consts.NacosContextPath),
			IpAddr:      localConfig.GetString(consts.NacosIpAddr),
			Port:        localConfig.GetUint64(consts.NacosPort),
		},
	}

	LocalConfig = localConfig
}
