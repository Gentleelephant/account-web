package server

import (
	"account-web/config"
	"account-web/initialize"
	"account-web/v1/handler"
	"fmt"
	"github.com/Gentleelephant/common/utils"
	"github.com/gin-gonic/gin"
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
	engine.Use(initialize.GinLogger())

	group := engine.Group("/v1")
	group.GET("/account/list", handler.GetAccountList)
	//router.RegisterGetRouter(engine, "/v1", handler.GetHandlerMap)

	// TODO 注册服务

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
	zap.L().Info("Server Stop", zap.String("signal", s.String()))
}
