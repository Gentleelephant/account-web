package server

import (
	"account-web/config"
	"account-web/initialize"
	"account-web/v1/handler"
	"fmt"
	"github.com/Gentleelephant/common/utils"
	"github.com/gin-gonic/gin"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func Start() {
	// 初始化配置
	config.GetLocalConfig()
	// 初始化日志配置
	initialize.InitLogger()
	// 初始化grpc客户端
	err := initialize.InitGrpcClient()
	if err != nil {
		panic(err)
	}
	ip, err := utils.GetOutBoundIP()
	if err != nil {
		ip, err = utils.GetInterfaceIpv4Addr("eth0")
		if err != nil {
			panic(err)
		}
	}
	port := config.LocalConfig.GetInt("service.port")
	if port == 0 {
		port = 8901
	}

	// 启动gin
	engine := gin.Default()
	engine.Use(gin.Recovery())
	//engine.Use(initialize.GinLogger())

	group := engine.Group("/v1")
	group.GET("/account/list", handler.GetAccountList)
	group.GET("/account/mobile", handler.GetAccountByMobile)
	group.GET("/account/id", handler.GetAccountByID)
	//router.RegisterGetRouter(engine, "/v1", handler.GetHandlerMap)

	// TODO 注册服务
	param := vo.RegisterInstanceParam{
		Ip:          ip,
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
	zap.L().Info("RegisterInstance", zap.Any("param", param))
	// TODO: 注册服务不成功
	ok, err := utils.RegisterInstance(*config.NacosConfig, param)
	if !ok {
		zap.L().Error("RegisterInstance failed")
		os.Exit(1)
	}
	if err != nil {
		zap.L().Error("RegisterInstance Error", zap.Error(err))
	}
	go func() {
		err := engine.Run(fmt.Sprintf("%s:%d", ip, port))
		if err != nil {
			fmt.Println("server error")
		}
	}()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	s := <-sig
	// TODO 注销服务
	deparams := vo.DeregisterInstanceParam{
		Ip:          ip,
		Port:        uint64(port),
		Cluster:     config.LocalConfig.GetString("nacos.cluster"),
		ServiceName: config.LocalConfig.GetString("service.name"),
		GroupName:   config.LocalConfig.GetString(config.LocalConfig.GetString("nacos.group")),
		Ephemeral:   true,
	}
	_, err = utils.DeregisterInstance(*config.NacosConfig, deparams)
	if err != nil {
		zap.L().Error("DeregisterInstance Error", zap.Error(err))
	}
	zap.L().Info("Server Stop", zap.String("signal", s.String()))
}
