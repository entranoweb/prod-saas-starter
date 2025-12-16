package cmd

import (
	"log"

	"github.com/joho/godotenv"
	"go.uber.org/dig"

	server "github.com/moasq/go-b2b-starter/server/domain"
)

func Execute() {
	// Load environment variables from app.env file
	if err := godotenv.Load("app.env"); err != nil {
		log.Printf("Warning: Error loading app.env file: %v", err)
	}

	container := dig.New()

	InitMods(container)

	var srv server.Server

	if err := container.Invoke(func(s server.Server) {
		srv = s
	}); err != nil {
		panic(err)
	}

	srv.Start()

}
