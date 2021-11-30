package app

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/defaulterrr/iot3/processing/internal/config"
	"github.com/defaulterrr/iot3/processing/internal/model"
	"github.com/defaulterrr/iot3/processing/internal/repository"
	"github.com/defaulterrr/iot3/processing/internal/service"
)

func Run() {
	// conn, err := grpc.Dial(":8082", grpc.WithInsecure())
	// if err != nil {
	// 	fmt.Printf("grpc.Dial: %v", err)
	// }
	// defer conn.Close()

	// client := pb.NewDHTClient(conn)

	// ctx, cancel := context.WithCancel(context.Background())

	// stream, err := client.GetDHTMetrics(ctx, &emptypb.Empty{})
	// if err != nil {
	// 	fmt.Printf("client.GetDHTMetrics: %v", err)
	// }

	// for i := 0; i < 10; i++ {
	// 	metrics, err := stream.Recv()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatalf("stream.Recv: %v", err)
	// 	}

	// 	fmt.Println(metrics)

	// 	i++
	// }

	// cancel()

	var pathToConfig string

	flag.StringVar(&pathToConfig, "config", "./config.yaml", "Specify a path to config file")
	flag.Parse()

	config, err := config.NewConfig(pathToConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conns, err := repository.GetGRPCConns(&config.Grpc)
	if err != nil {
		log.Fatalf("couldn't connect to grpc servers: %v", err)
		os.Exit(1)
	}

	defer conns.DHTConn.Close()

	repo := repository.NewRepository(conns)
	serv := service.NewService(repo)

	// testing GetDHTMetrics function
	metrics := make(chan model.DHTMetrics)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		err = serv.GetDHTMetrics(ctx, metrics)
		if err != nil {
			log.Fatalf("GetDHTMetrics: %v", err)
			os.Exit(1)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println(<-metrics)
	}
	cancel()

	fmt.Println("finish")
}
