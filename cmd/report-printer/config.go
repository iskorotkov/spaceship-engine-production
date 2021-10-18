package main

type Config struct {
	Port  int
	Input IOEntry
}

type IOEntry struct {
	Nats Server
}

type Server struct {
	Address string
}
