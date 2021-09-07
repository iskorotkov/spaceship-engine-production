package main

import (
	"github.com/iskorotkov/spaceship-engine-production/internal/generate"
)

type Config struct {
	Output IOEntry
	Data   Data
}

type IOEntry struct {
	ConnString ConnStrings
	Excel      Excel
}

type ConnStrings struct {
	SQLite     string
	PostgreSQL string
}

type Excel struct {
	File string
}

type Data struct {
	Seed       int
	Links      Links
	Clients    GeneratedHumanNames
	Providers  GeneratedHumanNames
	Components GeneratedHumanNames
	Engines    GeneratedDeviceNames
}

type Links struct {
	ClientOrders       Link
	EngineComponents   Link
	ComponentProviders Link
}

type Link struct {
	Count generate.Range
}

type GeneratedHumanNames struct {
	Count int
	Names HumanNames
}

type HumanNames struct {
	Adjectives []string
	Nouns      []string
}

type GeneratedDeviceNames struct {
	Count int
	Names DeviceNames
}

type DeviceNames struct {
	Types       []string
	Generations []string
	Series      []string
}
