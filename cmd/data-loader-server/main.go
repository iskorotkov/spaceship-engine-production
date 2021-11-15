package main

import (
	logpkg "log"
	"os"
	"os/signal"

	"github.com/iskorotkov/spaceship-engine-production/internal/models"
	"github.com/iskorotkov/spaceship-engine-production/internal/transport"
	"github.com/iskorotkov/spaceship-engine-production/internal/transport/nats"
	"github.com/iskorotkov/spaceship-engine-production/internal/transport/quic"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var log = logpkg.New(os.Stderr, "[main] ", logpkg.LstdFlags|logpkg.Lmsgprefix|logpkg.Lshortfile)

func main() {
	config := readConfig()
	// db := setupDB(config.ConnString)

	servers := map[string]transport.Server{
		config.AddrQUIC: &quic.Server{},
		config.AddrNATS: &nats.Server{},
	}

	for addr, s := range servers {
		if err := s.Start(addr); err != nil {
			log.Fatalf("error starting server: %v", err)
		}

		log.Printf("server started at %s", addr)

		if err := s.Handle("ping", func(req interface{}) (transport.Response, error) {
			return transport.Response{
				Message: "hello there",
			}, nil
		}); err != nil {
			log.Fatalf("error adding handler: %w", err)
		}
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	<-done
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
