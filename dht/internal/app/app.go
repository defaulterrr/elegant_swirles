package app

import (
	"flag"
	"fmt"
	"os"

	"github.com/defaulterrr/iot3/dht/internal/config"
	"github.com/defaulterrr/iot3/dht/internal/server"
	"github.com/defaulterrr/iot3/dht/internal/service"
)

func Run() {
	var pathToConfig string

	flag.StringVar(&pathToConfig, "config", "./config.yaml", "Specify a path to config file")
	flag.Parse()

	config, err := config.NewConfig(pathToConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	serv := service.NewService()

	if err := server.NewServer(serv).Start(&config.Grpc); err != nil {
		fmt.Printf("server.NewServer: %v", err)
		return
	}
}
