package main

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	// Setup the mgm default config
	err := mgm.SetDefaultConfig(nil, "notar-e-anotar", options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
	}
}
