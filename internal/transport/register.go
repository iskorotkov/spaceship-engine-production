package transport

import (
	"encoding/gob"

	"github.com/iskorotkov/spaceship-engine-production/internal/models"
)

//nolint:exhaustivestruct
func RegisterModels() {
	gob.Register(models.Client{})
	gob.Register(models.Component{})
	gob.Register(models.Engine{})
	gob.Register(models.Order{})
	gob.Register(models.Provider{})
	gob.Register(models.Requirement{})
	gob.Register(models.Supply{})
}
