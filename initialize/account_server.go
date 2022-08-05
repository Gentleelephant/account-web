package initialize

import (
	"account-web/config"
	"fmt"
	"github.com/Gentleelephant/common/utils"
	pb "github.com/Gentleelephant/proto-center/pb/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var AccountServiceClient pb.AccountServiceClient

func InitGrpcClient() error {

	host := config.LocalConfig.GetString("account-srv.host")
	port := config.LocalConfig.GetInt("account-srv.port")
	if host == "" && port == 0 {
		// 如果没有配置account-srv，则使用nacos中的配置
		instance, err := utils.SelectOneHealthyInstance(*config.NacosConfig, vo.SelectOneHealthInstanceParam{
			Clusters:    []string{config.LocalConfig.GetString("nacos.cluster")},
			ServiceName: config.LocalConfig.GetString("account-srv.name"),
			GroupName:   config.LocalConfig.GetString("nacos.group"),
		})
		if err != nil {
			zap.L().Error("SelectOneHealthyInstance Error", zap.Error(err))
		}
		host = instance.Ip
		port = int(instance.Port)
	}

	dial, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	AccountServiceClient = pb.NewAccountServiceClient(dial)
	return err
}

func RegisterWebService(host string, port int) {
	param := vo.RegisterInstanceParam{
		Ip:          host,
		Port:        uint64(port),
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Metadata:    nil,
		ClusterName: config.LocalConfig.GetString("nacos.cluster"),
		ServiceName: config.LocalConfig.GetString("service.name"),
		GroupName:   config.LocalConfig.GetString("nacos.group"),
		Ephemeral:   true,
	}
	ok, err := utils.RegisterInstance(*config.NacosConfig, param)
	if !ok {
		zap.L().Error("RegisterInstance failed")
	}
	if err != nil {
		zap.L().Error("RegisterInstance Error", zap.Error(err))
	}
}

func DeregisterWebService(host string, port int) {
	deparams := vo.DeregisterInstanceParam{
		Ip:          host,
		Port:        uint64(port),
		Cluster:     config.LocalConfig.GetString("nacos.cluster"),
		ServiceName: config.LocalConfig.GetString("service.name"),
		GroupName:   config.LocalConfig.GetString(config.LocalConfig.GetString("nacos.group")),
		Ephemeral:   true,
	}
	_, err := utils.DeregisterInstance(*config.NacosConfig, deparams)
	if err != nil {
		zap.L().Error("DeregisterInstance Error", zap.Error(err))
	}
}
