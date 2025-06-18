package main

import (
	"log"

	"github.com/Pratam-Kalligudda/order-service-go/config"
	"github.com/Pratam-Kalligudda/order-service-go/internal/api"
)

func main() {
	config, err := config.SetupEnv()
	if err != nil {
		log.Fatalf("err : %v", err.Error())
	}

	api.SetupServer(config)
}
