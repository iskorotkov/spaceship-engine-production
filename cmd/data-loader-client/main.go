package main

import (
	"crypto/tls"
	"log"

	"github.com/iskorotkov/spaceship-engine-production/internal/models"
	"github.com/iskorotkov/spaceship-engine-production/internal/transports/quic"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	config := readConfig()
	// db := setupDB(config)

	client, err := quic.New(config.Addr, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	if err := client.Send("ping", "hi"); err != nil {
		log.Fatal(err)
	}

	resp, err := client.Recv()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("received %q", resp)
}

func normalizeSupplies(inputDB *gorm.DB, outputDB *gorm.DB, batchSize int) {
	var rowSupplies []struct {
		ProviderID, ComponentID uint
	}

	inputDB.
		Model(&models.Row{}). //nolint:exhaustivestruct
		Distinct("provider_id, component_id").
		Find(&rowSupplies)

	supplies := make([]models.Supply, 0, len(rowSupplies))

	for i, s := range rowSupplies {
		supplies = append(supplies, models.Supply{
			ID:          uint(i + 1),
			Component:   nil,
			ComponentID: s.ComponentID,
			Provider:    nil,
			ProviderID:  s.ProviderID,
		})
	}

	outputDB.CreateInBatches(&supplies, batchSize)
}

func normalizeRequirements(inputDB *gorm.DB, outputDB *gorm.DB, batchSize int) {
	var rowRequirements []struct {
		EngineID, ComponentID uint
	}

	inputDB.
		Model(&models.Row{}). //nolint:exhaustivestruct
		Distinct("engine_id, component_id").
		Find(&rowRequirements)

	requirements := make([]models.Requirement, 0, len(rowRequirements))

	for i, r := range rowRequirements {
		requirements = append(requirements, models.Requirement{
			ID:          uint(i + 1),
			Engine:      nil,
			EngineID:    r.EngineID,
			Component:   nil,
			ComponentID: r.ComponentID,
		})
	}

	outputDB.CreateInBatches(&requirements, batchSize)
}

func normalizeOrders(inputDB *gorm.DB, outputDB *gorm.DB, batchSize int) {
	var rowOrders []struct {
		models.RowOrder
		models.RowClient
		models.RowEngine
	}

	inputDB.
		Model(&models.Row{}). //nolint:exhaustivestruct
		Distinct("order_id, order_amount, order_created_at, order_completed_at, client_id, engine_id").
		Find(&rowOrders)

	orders := make([]models.Order, 0, len(rowOrders))

	for _, o := range rowOrders {
		orders = append(orders, models.Order{
			ID:          o.OrderID,
			Amount:      o.OrderAmount,
			CreatedAt:   o.OrderCreatedAt,
			CompletedAt: o.OrderCompletedAt,
			Client:      nil,
			ClientID:    o.ClientID,
			Engine:      nil,
			EngineID:    o.EngineID,
		})
	}

	outputDB.CreateInBatches(&orders, batchSize)
}

func normalizeProviders(inputDB *gorm.DB, outputDB *gorm.DB, batchSize int) {
	var rowProviders []models.RowProvider

	inputDB.
		Model(&models.Row{}). //nolint:exhaustivestruct
		Distinct("provider_id, provider_name").
		Find(&rowProviders)

	providers := make([]models.Provider, 0, len(rowProviders))

	for _, p := range rowProviders {
		providers = append(providers, models.Provider{
			ID:   p.ProviderID,
			Name: p.ProviderName,
		})
	}

	outputDB.CreateInBatches(&providers, batchSize)
}

func normalizeClients(inputDB *gorm.DB, outputDB *gorm.DB, batchSize int) {
	var rowClients []models.RowClient

	inputDB.
		Model(&models.Row{}). //nolint:exhaustivestruct
		Distinct("client_id, client_name").
		Find(&rowClients)

	clients := make([]models.Client, 0, len(rowClients))

	for _, c := range rowClients {
		clients = append(clients, models.Client{
			ID:   c.ClientID,
			Name: c.ClientName,
		})
	}

	outputDB.CreateInBatches(&clients, batchSize)
}

func normalizeComponents(inputDB *gorm.DB, outputDB *gorm.DB, batchSize int) {
	var rowComponents []models.RowComponent

	inputDB.
		Model(&models.Row{}). //nolint:exhaustivestruct
		Distinct("component_id, component_name").
		Find(&rowComponents)

	components := make([]models.Component, 0, len(rowComponents))

	for _, c := range rowComponents {
		components = append(components, models.Component{
			ID:   c.ComponentID,
			Name: c.ComponentName,
		})
	}

	outputDB.CreateInBatches(&components, batchSize)
}

func normalizeEngines(inputDB *gorm.DB, outputDB *gorm.DB, batchSize int) {
	var rowEngines []models.RowEngine

	inputDB.
		Model(&models.Row{}). //nolint:exhaustivestruct
		Distinct("engine_id, engine_name, engine_power").
		Find(&rowEngines)

	engines := make([]models.Engine, 0, len(rowEngines))

	for _, e := range rowEngines {
		engines = append(engines, models.Engine{
			ID:    e.EngineID,
			Name:  e.EngineName,
			Power: e.EnginePower,
		})
	}

	outputDB.CreateInBatches(&engines, batchSize)
}

func setupDB(config Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(config.ConnString))
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
	viper.AddConfigPath("./config/data-loader-client")
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
