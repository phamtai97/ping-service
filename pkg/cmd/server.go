package cmd

import (
	"context"
	"pingservice/configs"
	"pingservice/pkg/core"
	grpc "pingservice/pkg/protocol/grpc"
	pingservice "pingservice/pkg/service"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//RunServer run gRPC server
func RunServer() error {
	ctx := context.Background()

	//load config
	config := &configs.PingServiceConfig{}
	configs.LoadConfig()
	if err := viper.Unmarshal(config); err != nil {
		log.Fatal("load config: ", err)
	}

	pingCore, err := core.NewPingCore()
	if err != nil {
		return err
	}

	pingService := pingservice.NewPingService(pingCore)

	return grpc.RunServer(ctx, pingService, strconv.Itoa(config.GRPCPort))
}
