package main

type Config struct {
	Input  IOEntry
	Output IOEntry
}

type IOEntry struct {
	ConnString ConnStrings
}

type ConnStrings struct {
	SQLite     string
	PostgreSQL string
}
