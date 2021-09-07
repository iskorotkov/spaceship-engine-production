package main

import (
	"log"
	"net"

	"github.com/iskorotkov/spaceship-engine-production/api"
	"github.com/iskorotkov/spaceship-engine-production/internal/printer"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("error starting tcp listener: %v", err)
	}

	server := grpc.NewServer()
	api.RegisterReportPrinterServer(server, printer.NewServer())

	log.Fatal(server.Serve(lis))
}
