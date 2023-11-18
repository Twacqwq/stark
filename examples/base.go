package main

import (
	"log"

	"github.com/FarmerChillax/stark"
	"github.com/FarmerChillax/stark/app"
	starkConf "github.com/FarmerChillax/stark/config"
)

func main() {
	app, err := app.New(&stark.Application{
		Name:   "base-demo",
		Host:   "127.0.0.1",
		Port:   6000,
		Config: &starkConf.Config{},
		LoadConfig: func() error {
			return nil
		},
		SetupVars: func() error {
			return nil
		},
		RegisterCallback: make(map[stark.CallbackPosition]stark.CallbackFunc),
		// RegisterRouter
	})
	if err != nil {
		log.Fatalf("app.New err: %v", err)
	}

	if err := app.ListenAndServe(); err != nil {
		log.Fatalf("app.ListenAndServe err: %v", err)
	}
}
