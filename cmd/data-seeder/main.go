package main

import (
	"log"
	"math/rand"

	"github.com/iskorotkov/spaceship-engine-production/internal/generate"
	"github.com/iskorotkov/spaceship-engine-production/internal/models"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const batchSize = 100

func main() {
	config := readConfig()
	db := setupDB(config)

	rand.Seed(int64(config.Data.Seed))

	components := generate.Components(
		config.Data.Components.Names.Adjectives,
		config.Data.Components.Names.Nouns,
		config.Data.Components.Count,
	)

	engines := generate.Engines(
		config.Data.Engines.Names.Types,
		config.Data.Engines.Names.Generations,
		config.Data.Engines.Names.Series,
		config.Data.Engines.Count,
	)

	clients := generate.Clients(
		config.Data.Clients.Names.Adjectives,
		config.Data.Clients.Names.Nouns,
		config.Data.Clients.Count,
	)

	providers := generate.Providers(
		config.Data.Providers.Names.Adjectives,
		config.Data.Providers.Names.Nouns,
		config.Data.Providers.Count,
	)

	rows := generate.Rows(
		clients,
		engines,
		components,
		providers,
		config.Data.Links.ClientOrders.Count,
		config.Data.Links.EngineComponents.Count,
		config.Data.Links.ComponentProviders.Count,
	)

	db.Delete(&models.Row{}, "true") //nolint:exhaustivestruct
	db.CreateInBatches(&rows, batchSize)
}

func setupDB(config Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(config.Output.ConnString.SQLite))
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
	viper.AddConfigPath("./config/data-seeder")
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
