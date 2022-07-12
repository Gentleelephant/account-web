package server

import (
	"account-web/config"
	"account-web/initialize"
	"account-web/v1/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

func Start() error {
	// 初始化配置
	config.LoadYaml()

	// 初始化grpc客户端
	err := initialize.InitGrpcClient()
	if err != nil {
		return err
	}

	// 启动gin
	engine := gin.Default()
	engine.Use(gin.Recovery())

	group := engine.Group("/account")
	{
		group.GET("list", handler.GetAccountList)
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := engine.Run(":8080")
		if err != nil {
			fmt.Println("server error")
		}
	}()
	wg.Wait()
	return nil
}
