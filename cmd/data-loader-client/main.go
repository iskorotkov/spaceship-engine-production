package main

import (
	"context"
	logpkg "log"
	"os"

	data_loader "github.com/iskorotkov/spaceship-engine-production/api/data-loader"
	"github.com/iskorotkov/spaceship-engine-production/internal/models"
	"github.com/iskorotkov/spaceship-engine-production/internal/transport"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var log = logpkg.New(os.Stderr, "[main] ", logpkg.LstdFlags|logpkg.Lmsgprefix|logpkg.Lshortfile)

func main() {
	config := readConfig()
	db := setupDB(config.ConnString)

	transport.RegisterModels()

	tlsConfig := transport.MustCreateTLSConfig(config.CertFile, config.KeyFile, config.RootCA)
	tlsConfig.InsecureSkipVerify = true

	conn, err := grpc.Dial(config.AddrGRPC, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
	if err != nil {
		log.Fatalf("error creating grpc connection: %v", err)
	}

	client := data_loader.NewDataLoaderClient(conn)

	for _, value := range must(normalizeClients(db)).([]models.Client) {
		log.Printf("sending client %d", value.ID)

		if _, err := client.SaveClient(context.Background(), &data_loader.ClientReq{
			Id:   uint32(value.ID),
			Name: value.Name,
		}); err != nil {
			log.Fatalf("error saving client: %v", err)
		}
	}

	for _, value := range must(normalizeComponents(db)).([]models.Component) {
		log.Printf("sending component %d", value.ID)

		if _, err := client.SaveComponent(context.Background(), &data_loader.ComponentReq{
			Id:   uint32(value.ID),
			Name: value.Name,
		}); err != nil {
			log.Fatalf("error saving component: %v", err)
		}
	}

	for _, value := range must(normalizeEngines(db)).([]models.Engine) {
		log.Printf("sending engine %d", value.ID)

		if _, err := client.SaveEngine(context.Background(), &data_loader.EngineReq{
			Id:    uint32(value.ID),
			Name:  value.Name,
			Power: value.Power,
		}); err != nil {
			log.Fatalf("error saving engine: %v", err)
		}
	}

	for _, value := range must(normalizeOrders(db)).([]models.Order) {
		log.Printf("sending order %d", value.ID)

		if _, err := client.SaveOrder(context.Background(), &data_loader.OrderReq{
			Id:          uint32(value.ID),
			Amount:      int32(value.Amount),
			CreatedAt:   value.CreatedAt.UnixMilli(),
			CompletedAt: value.CompletedAt.UnixMilli(),
			ClientId:    uint32(value.ClientID),
			EngineId:    uint32(value.EngineID),
		}); err != nil {
			log.Fatalf("error saving order: %v", err)
		}
	}

	for _, value := range must(normalizeProviders(db)).([]models.Provider) {
		log.Printf("sending provider %d", value.ID)

		if _, err := client.SaveProvider(context.Background(), &data_loader.ProviderReq{
			Id:   uint32(value.ID),
			Name: value.Name,
		}); err != nil {
			log.Fatalf("error saving provider: %v", err)
		}
	}

	for _, value := range must(normalizeRequirements(db)).([]models.Requirement) {
		log.Printf("sending requirement %d", value.ID)

		if _, err := client.SaveRequirement(context.Background(), &data_loader.RequirementReq{
			Id:          uint32(value.ID),
			EngineId:    uint32(value.EngineID),
			ComponentId: uint32(value.ComponentID),
		}); err != nil {
			log.Fatalf("error saving requirement: %v", err)
		}
	}

	for _, value := range must(normalizeSupplies(db)).([]models.Supply) {
		log.Printf("sending supply %d", value.ID)

		if _, err := client.SaveSupply(context.Background(), &data_loader.SupplyReq{
			Id:          uint32(value.ID),
			ComponentId: uint32(value.ComponentID),
			ProviderId:  uint32(value.ProviderID),
		}); err != nil {
			log.Fatalf("error saving supply: %v", err)
		}
	}
}

func must(v interface{}, err error) interface{} {
	if err != nil {
		log.Fatalf("error is not nil: %v", err)
	}

	return v
}

func normalizeSupplies(inputDB *gorm.DB) ([]models.Supply, error) {
	var rowSupplies []struct {
		ProviderID, ComponentID uint
	}

	if err := inputDB.
		Model(&models.Row{}). //nolint:exhaustivestruct
		Distinct("provider_id, component_id").
		Find(&rowSupplies).Error; err != nil {
		return nil, err
	}

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

	return supplies, nil
}

func normalizeRequirements(inputDB *gorm.DB) ([]models.Requirement, error) {
	var rowRequirements []struct {
		EngineID, ComponentID uint
	}

	if err := inputDB.
		Model(&models.Row{}). //nolint:exhaustivestruct
		Distinct("engine_id, component_id").
		Find(&rowRequirements).Error; err != nil {
		return nil, err
	}

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

	return requirements, nil
}

func normalizeOrders(inputDB *gorm.DB) ([]models.Order, error) {
	var rowOrders []struct {
		models.RowOrder
		models.RowClient
		models.RowEngine
	}

	if err := inputDB.
		Model(&models.Row{}). //nolint:exhaustivestruct
		Distinct("order_id, order_amount, order_created_at, order_completed_at, client_id, engine_id").
		Find(&rowOrders).Error; err != nil {
		return nil, err
	}

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

	return orders, nil
}

func normalizeProviders(inputDB *gorm.DB) ([]models.Provider, error) {
	var rowProviders []models.RowProvider

	if err := inputDB.
		Model(&models.Row{}). //nolint:exhaustivestruct
		Distinct("provider_id, provider_name").
		Find(&rowProviders).Error; err != nil {
		return nil, err
	}

	providers := make([]models.Provider, 0, len(rowProviders))

	for _, p := range rowProviders {
		providers = append(providers, models.Provider{
			ID:   p.ProviderID,
			Name: p.ProviderName,
		})
	}

	return providers, nil
}

func normalizeClients(inputDB *gorm.DB) ([]models.Client, error) {
	var rowClients []models.RowClient

	if err := inputDB.
		Model(&models.Row{}). //nolint:exhaustivestruct
		Distinct("client_id, client_name").
		Find(&rowClients).Error; err != nil {
		return nil, err
	}

	clients := make([]models.Client, 0, len(rowClients))

	for _, c := range rowClients {
		clients = append(clients, models.Client{
			ID:   c.ClientID,
			Name: c.ClientName,
		})
	}

	return clients, nil
}

func normalizeComponents(inputDB *gorm.DB) ([]models.Component, error) {
	var rowComponents []models.RowComponent

	if err := inputDB.
		Model(&models.Row{}). //nolint:exhaustivestruct
		Distinct("component_id, component_name").
		Find(&rowComponents).Error; err != nil {
		return nil, err
	}

	components := make([]models.Component, 0, len(rowComponents))

	for _, c := range rowComponents {
		components = append(components, models.Component{
			ID:   c.ComponentID,
			Name: c.ComponentName,
		})
	}

	return components, nil
}

func normalizeEngines(inputDB *gorm.DB) ([]models.Engine, error) {
	var rowEngines []models.RowEngine

	if err := inputDB.
		Model(&models.Row{}). //nolint:exhaustivestruct
		Distinct("engine_id, engine_name, engine_power").
		Find(&rowEngines).Error; err != nil {
		return nil, err
	}

	engines := make([]models.Engine, 0, len(rowEngines))

	for _, e := range rowEngines {
		engines = append(engines, models.Engine{
			ID:    e.EngineID,
			Name:  e.EngineName,
			Power: e.EnginePower,
		})
	}

	return engines, nil
}

func setupDB(connString string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(connString))
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
	viper.AddConfigPath("/config")
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
