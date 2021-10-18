package main

const (
	ClientGRPC Client = "grpc"
	ClientNATS Client = "nats"
)

type Client string

type Config struct {
	Client Client
	Input  IOEntry
	Output IOEntry
}

type IOEntry struct {
	Server     Server
	Nats       Server
	ConnString ConnStrings
	Excel      Excel
}

type Server struct {
	Address string
}

type ConnStrings struct {
	SQLite     string
	PostgreSQL string
}

type Excel struct {
	File string
}
