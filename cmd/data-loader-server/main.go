package main

import (
	logpkg "log"
	"os"
	"os/signal"

	"github.com/iskorotkov/spaceship-engine-production/internal/models"
	"github.com/iskorotkov/spaceship-engine-production/internal/transport"
	"github.com/iskorotkov/spaceship-engine-production/internal/transport/grpc"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var log = logpkg.New(os.Stderr, "[main] ", logpkg.LstdFlags|logpkg.Lmsgprefix|logpkg.Lshortfile)

func main() {
	config := readConfig()
	db := setupDB(config.ConnString)

	tlsConfig := transport.MustCreateTLSConfig(config.CertFile, config.KeyFile, config.RootCA)

	server := grpc.NewServer(db)

	go func() {
		if err := server.Start(config.AddrGRPC, tlsConfig); err != nil {
			log.Printf("error in server: %v", err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	<-done

	if err := server.Close(); err != nil {
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
