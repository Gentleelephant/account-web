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
	host := config.LocalConfig.GetString("account-srv.host")
	port := config.LocalConfig.GetInt("account-srv.port")

	dial, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	AccountServiceClient = pb.NewAccountServiceClient(dial)
	return err
}
