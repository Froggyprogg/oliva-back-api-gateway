package config

import "github.com/spf13/viper"

type Config struct {
	Port
	AuthSrv
	ChatSrv
	MainSrv
}
type Port struct {
	Port string
}
type AuthSrv struct {
	Host string
	Port string
}
type ChatSrv struct {
	Host string
	Port string
}
type MainSrv struct {
	Host string
	Port string
}

func NewConfig() *Config {
	return &Config{
		Port: Port{
			Port: viper.GetString("port.port"),
		},
		AuthSrv: AuthSrv{
			Port: viper.GetString("auth-srv.port"),
			Host: viper.GetString("auth-srv.host"),
		},
		ChatSrv: ChatSrv{
			Port: viper.GetString("chat-srv.port"),
			Host: viper.GetString("chat-srv.host"),
		},
		MainSrv: MainSrv{
			Port: viper.GetString("main-srv.port"),
			Host: viper.GetString("main-srv.host"),
		},
	}
}

func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
