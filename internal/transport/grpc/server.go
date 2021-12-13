package grpc

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"time"

	data_loader "github.com/iskorotkov/spaceship-engine-production/api/data-loader"
	"github.com/iskorotkov/spaceship-engine-production/internal/models"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func savingError(err error) error {
	return fmt.Errorf("error saving model: %w", err)
}

var _ = (data_loader.DataLoaderServer)(&Server{})

type Server struct {
	data_loader.UnsafeDataLoaderServer
	db     *gorm.DB
	conn   net.Listener
	server *grpc.Server
}

func NewServer(db *gorm.DB) *Server {
	return &Server{
		UnsafeDataLoaderServer: nil,
		db:                     db,
		conn:                   nil,
		server:                 nil,
	}
}

func (s Server) SaveClient(_ context.Context, req *data_loader.ClientReq) (*data_loader.ClientResp, error) {
	log.Printf("saving client %d", req.Id)

	if err := s.db.Save(&models.Client{
		ID:   uint(req.Id),
		Name: req.Name,
	}).Error; err != nil {
		return nil, savingError(err)
	}

	return &data_loader.ClientResp{}, nil
}

func (s Server) SaveComponent(_ context.Context, req *data_loader.ComponentReq) (*data_loader.ComponentResp, error) {
	log.Printf("saving component %d", req.Id)

	if err := s.db.Save(&models.Component{
		ID:   uint(req.Id),
		Name: req.Name,
	}).Error; err != nil {
		return nil, savingError(err)
	}

	return &data_loader.ComponentResp{}, nil
}

func (s Server) SaveEngine(_ context.Context, req *data_loader.EngineReq) (*data_loader.EngineResp, error) {
	log.Printf("saving engine %d", req.Id)

	if err := s.db.Save(&models.Engine{
		ID:    uint(req.Id),
		Name:  req.Name,
		Power: req.Power,
	}).Error; err != nil {
		return nil, savingError(err)
	}

	return &data_loader.EngineResp{}, nil
}

func (s Server) SaveOrder(_ context.Context, req *data_loader.OrderReq) (*data_loader.OrderResp, error) {
	log.Printf("saving order %d", req.Id)

	completedAt := time.UnixMilli(req.CompletedAt)

	if err := s.db.Save(&models.Order{
		ID:          uint(req.Id),
		Amount:      int(req.Amount),
		CreatedAt:   time.UnixMilli(req.CreatedAt),
		CompletedAt: &completedAt,
		ClientID:    uint(req.ClientId),
		EngineID:    uint(req.EngineId),
	}).Error; err != nil {
		return nil, savingError(err)
	}

	return &data_loader.OrderResp{}, nil
}

func (s Server) SaveProvider(_ context.Context, req *data_loader.ProviderReq) (*data_loader.ProviderResp, error) {
	log.Printf("saving provider %d", req.Id)

	if err := s.db.Save(&models.Provider{
		ID:   uint(req.Id),
		Name: req.Name,
	}).Error; err != nil {
		return nil, savingError(err)
	}

	return &data_loader.ProviderResp{}, nil
}

func (s Server) SaveRequirement(_ context.Context, req *data_loader.RequirementReq) (*data_loader.RequirementResp, error) {
	log.Printf("saving requirement %d", req.Id)

	if err := s.db.Save(&models.Requirement{
		ID:          uint(req.Id),
		EngineID:    uint(req.EngineId),
		ComponentID: uint(req.ComponentId),
	}).Error; err != nil {
		return nil, savingError(err)
	}

	return &data_loader.RequirementResp{}, nil
}

func (s Server) SaveSupply(_ context.Context, req *data_loader.SupplyReq) (*data_loader.SupplyResp, error) {
	log.Printf("saving supply %d", req.Id)

	if err := s.db.Save(&models.Supply{
		ID:          uint(req.Id),
		ComponentID: uint(req.ComponentId),
		ProviderID:  uint(req.ProviderId),
	}).Error; err != nil {
		return nil, savingError(err)
	}

	return &data_loader.SupplyResp{}, nil
}

func (s Server) Start(addr string, config *tls.Config) error {
	log.Printf("starting server on %s", addr)

	var err error

	s.conn, err = tls.Listen("tcp", addr, config)
	if err != nil {
		return fmt.Errorf("error listening on address %s: %w", addr, err)
	}

	s.server = grpc.NewServer()
	data_loader.RegisterDataLoaderServer(s.server, s)

	if err := s.server.Serve(s.conn); err != nil {
		return fmt.Errorf("error serving: %w", err)
	}

	return nil
}

func (s Server) Close() error {
	log.Printf("stopping server")
	s.server.GracefulStop()
	return s.Close()
}
