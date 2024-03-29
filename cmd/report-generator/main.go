package main

import (
	"context"
	"log"
	"time"

	"github.com/iskorotkov/spaceship-engine-production/api/report-printer"
	"github.com/iskorotkov/spaceship-engine-production/internal/models"
	"github.com/iskorotkov/spaceship-engine-production/internal/printer"
	"github.com/iskorotkov/spaceship-engine-production/internal/printer/grpc"
	"github.com/iskorotkov/spaceship-engine-production/internal/printer/nats"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	topListSize = 10
	timeout     = time.Second * 10
)

func main() {
	config := readConfig()
	db := setupDB(config)

	log.Printf("generating report")

	totalOrders := getTotalOrders(db)
	topClients := getTopClients(db)
	topEngines := getTopEngines(db)

	req := createRequest(totalOrders, topClients, topEngines, config)

	log.Printf("creating a client based on config value %q", config.Client)

	var (
		client printer.Client
		err    error
	)

	switch config.Client {
	case ClientGRPC:
		client, err = grpc.NewClient(config.Output.Server.Address)
	case ClientNATS:
		client, err = nats.NewClient(config.Output.Nats.Address)
	default:
		log.Fatalf("invalid client in config file: client %q not supported", config.Client)
	}

	if err != nil {
		log.Fatalf("error creating printer client: %v", err)
	}

	log.Printf("sending report print request")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err = client.Print(ctx, req); err != nil {
		log.Printf("error processing print request: %v", err)

		return
	}
}

func createRequest(totalOrders int64, topClients []clientDTO, topEngines []engineDTO, config Config) *report_printer.PrintRequest {
	reqClients := make([]*report_printer.Client, 0, len(topClients))

	for _, c := range topClients {
		reqClients = append(reqClients, &report_printer.Client{
			Name:   c.Name,
			Orders: int32(c.OrdersCount),
		})
	}

	reqEngines := make([]*report_printer.Engine, 0, len(topEngines))

	for _, e := range topEngines {
		reqEngines = append(reqEngines, &report_printer.Engine{
			Name:   e.Name,
			Power:  e.Power,
			Orders: int32(e.OrdersCount),
		})
	}

	req := &report_printer.PrintRequest{
		Filepath:    config.Output.Excel.File,
		TotalOrders: int32(totalOrders),
		TopClients:  reqClients,
		TopEngines:  reqEngines,
	}

	return req
}

type engineDTO struct {
	models.Engine
	OrdersCount int64
}

func getTopEngines(db *gorm.DB) []engineDTO {
	var topEngines []engineDTO

	db.
		Model(&models.Engine{}). //nolint:exhaustivestruct
		Select("engines.id, engines.name, engines.power, count(*) as orders_count").
		Joins("JOIN orders ON orders.engine_id = engines.id").
		Group("engines.id").
		Order("orders_count DESC").
		Limit(topListSize).
		Find(&topEngines)

	return topEngines
}

type clientDTO struct {
	models.Client
	OrdersCount int64
}

func getTopClients(db *gorm.DB) []clientDTO {
	var topClients []clientDTO

	db.
		Model(&models.Client{}). //nolint:exhaustivestruct
		Select("clients.id, clients.name, count(*) as orders_count").
		Joins("JOIN orders ON orders.client_id = clients.id").
		Group("clients.id").
		Order("orders_count DESC").
		Limit(topListSize).
		Find(&topClients)

	return topClients
}

func getTotalOrders(db *gorm.DB) int64 {
	var totalOrders int64

	db.
		Model(models.Order{}). //nolint:exhaustivestruct
		Count(&totalOrders)

	return totalOrders
}

func setupDB(config Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.Input.ConnString.PostgreSQL))
	if err != nil {
		log.Fatalf("error opening DB: %v", err)
	}

	err = db.AutoMigrate(&models.Row{}) //nolint:exhaustivestruct
	if err != nil {
		log.Fatalf("error running migrations: %v", err)
	}

	return db
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
