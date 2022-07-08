package config

type AccountConfig struct {
	AccountServerConfig *AccountGrpcServerConfig `json:"account_server" yaml:"account_server"`
}

type AccountGrpcServerConfig struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}
