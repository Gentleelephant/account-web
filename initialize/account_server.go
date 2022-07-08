package initialize

import (
	"account-web/config"
	"fmt"
	pb "github.com/Gentleelephant/proto-center/pb/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var AccountServiceClient pb.AccountServiceClient

func InitGrpcClient() error {

	fmt.Println("Host:", config.Config.AccountServerConfig.Host)
	fmt.Println("Port:", config.Config.AccountServerConfig.Port)

	dial, err := grpc.Dial(fmt.Sprintf("%s:%d", config.Config.AccountServerConfig.Host, config.Config.AccountServerConfig.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	AccountServiceClient = pb.NewAccountServiceClient(dial)
	return err
}
