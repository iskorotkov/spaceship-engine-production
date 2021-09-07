package main

type Config struct {
	Input  IOEntry
	Output IOEntry
}

type IOEntry struct {
	Server     Server
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
