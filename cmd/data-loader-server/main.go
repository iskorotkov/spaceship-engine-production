package main

import (
	"fmt"
	logpkg "log"
	"os"
	"os/signal"
	"reflect"

	"github.com/iskorotkov/spaceship-engine-production/internal/models"
	"github.com/iskorotkov/spaceship-engine-production/internal/transport"
	"github.com/iskorotkov/spaceship-engine-production/internal/transport/nats"
	"github.com/iskorotkov/spaceship-engine-production/internal/transport/quic"
	"github.com/iskorotkov/spaceship-engine-production/internal/transport/tcp"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var log = logpkg.New(os.Stderr, "[main] ", logpkg.LstdFlags|logpkg.Lmsgprefix|logpkg.Lshortfile)

func main() {
	config := readConfig()
	db := setupDB(config.ConnString)

	transport.RegisterModels()

	tlsConfig := transport.MustCreateTLSConfig(config.CertFile, config.KeyFile, config.RootCA)

	servers := map[string]transport.Server{
		config.AddrTCP:  tcp.NewServer(),
		config.AddrQUIC: quic.NewServer(),
		config.AddrNATS: nats.NewServer(),
	}

	pingHandler := func(req interface{}) (transport.Response, error) {
		return transport.Response{Message: "hello there"}, nil
	}

	saveModelHandler := func(req interface{}) (transport.Response, error) {
		var err error

		switch req := req.(type) {
		case models.Client:
			err = db.Save(&req).Error
		case models.Component:
			err = db.Save(&req).Error
		case models.Engine:
			err = db.Save(&req).Error
		case models.Order:
			err = db.Save(&req).Error
		case models.Provider:
			err = db.Save(&req).Error
		case models.Requirement:
			err = db.Save(&req).Error
		case models.Supply:
			err = db.Save(&req).Error
		default:
			return transport.Response{}, fmt.Errorf("unsupported model type %s", reflect.TypeOf(req).Name())
		}

		if err != nil {
			return transport.Response{}, fmt.Errorf("error saving model: %w", err)
		}

		return transport.Response{Message: "success"}, nil
	}

	for addr, s := range servers {
		if err := s.Start(addr, tlsConfig); err != nil {
			log.Fatalf("error starting server: %v", err)
		}

		log.Printf("server started at %s", addr)

		if err := s.Handle("ping", pingHandler); err != nil {
			log.Fatalf("error adding handler: %v", err)
		}

		for _, model := range []string{"client", "component", "engine", "order", "provider", "requirement", "supply"} {
			if err := s.Handle(fmt.Sprintf("save/%s", model), saveModelHandler); err != nil {
				log.Fatalf("error adding handler: %v", err)
			}
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
