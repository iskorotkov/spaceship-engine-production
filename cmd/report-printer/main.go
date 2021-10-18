package main

import (
	"context"
	"log"
	"time"

	"github.com/iskorotkov/spaceship-engine-production/internal/printer/grpc"
	"github.com/iskorotkov/spaceship-engine-production/internal/printer/nats"
	"github.com/spf13/viper"
)

const timeout = time.Second * 10

func main() {
	config := readConfig()

	natsServer, err := nats.NewServer(config.Input.Nats.Address)
	if err != nil {
		log.Fatalf("error creating nats server: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := natsServer.Subscribe(ctx, nats.SubjectPrintRequests); err != nil {
		log.Printf("error subscribing to nats subject: %v", err)

		return
	}

	defer func(natsServer nats.Server) {
		_ = natsServer.Close()
	}(natsServer)

	grpcServer := grpc.NewServer(config.Port)
	if err := grpcServer.Listen(); err != nil {
		log.Print(err)

		return
	}
}

func readConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/report-generator")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	var config Config
	if err := viper.UnmarshalExact(&config); err != nil {
		log.Fatalf("error unmarshaling config: %v", err)
	}

	return config
}
