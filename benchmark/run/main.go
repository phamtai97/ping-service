package main

import (
	"pingservice/configs"

	log "github.com/sirupsen/logrus"

	"pingservice/benchmark"

	"github.com/spf13/viper"
)

func run() {
	//load config
	config := &configs.PingServiceConfig{}
	configs.LoadConfig()
	if err := viper.Unmarshal(config); err != nil {
		log.Fatal("load config: ", err)
	}

	managerClient := benchmark.NewManagerClient(config)

	boomerClient := benchmark.BoomerClient{}
	boomerClient.LoadManagerClient(managerClient)

	tasks, _ := boomerClient.LoadTask(benchmark.PING, 1)

	boomerClient.RunTask(tasks)
}

func main() {
	run()
}
