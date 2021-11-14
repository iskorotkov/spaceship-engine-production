package main

import (
	"log"

	"github.com/iskorotkov/spaceship-engine-production/internal/models"
	"github.com/iskorotkov/spaceship-engine-production/internal/transports/quic"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config := readConfig()
	// db := setupDB(config.ConnString)

	var server quic.Server
	server.Handle("ping", func(req interface{}) (quic.Response, error) {
		return quic.Response{
			Message: "hello there",
		}, nil
	})

	if err := server.Start(config.Addr); err != nil {
		log.Fatal(err)
	}
}

func setupDB(connString string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connString))
	if err != nil {
		log.Fatalf("error opening DB: %v", err)
	}

	//nolint:exhaustivestruct
	err = db.AutoMigrate(
		&models.Engine{},
		&models.Component{},
		&models.Client{},
		&models.Provider{},
		&models.Order{},
		&models.Requirement{},
		&models.Supply{},
	)

	if err != nil {
		log.Fatalf("error running migrations: %v", err)
	}

	return db
}

func readConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/data-loader-server")
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
