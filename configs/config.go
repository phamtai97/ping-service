package configs

import (
	"github.com/spf13/viper"
)

const configFilePath = "."
const configFileName = "config"

//PingServiceConfig struct
type PingServiceConfig struct {
	GRPCPort int
	GRPCHost string
	PoolSize int
	TimeOut  int
}

//LoadConfig load config
func LoadConfig() error {
	viper.SetConfigName(configFileName)
	viper.AddConfigPath(configFilePath)
	viper.SetConfigType("yaml")

	return viper.ReadInConfig()
}
