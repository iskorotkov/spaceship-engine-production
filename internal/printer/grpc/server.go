package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/iskorotkov/spaceship-engine-production/api"
	"github.com/iskorotkov/spaceship-engine-production/internal/printer"
	"google.golang.org/grpc"
)

type Server struct {
	port int
}

func NewServer(port int) Server {
	log.Printf("grpc server created")

	return Server{
		port: port,
	}
}

func (s Server) Listen() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("error starting tcp listener: %w", err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterReportPrinterServer(grpcServer, printer.New())

	log.Printf("grpc server started listening on port %d", s.port)

	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("error in grpc server: %w", err)
	}

	return nil
}
