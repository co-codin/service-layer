package main

import (
	"context"
	"fmt"
	"log"

	"github.com/co-codin/service-layer/internal/db"
	"github.com/joho/godotenv"
)

func Run() error {
	fmt.Println("starting up application")

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}

	db, err := db.NewDatabase()

	if err != nil {
		fmt.Println("Failed to connect to database")
		return err
	}

	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("rest api")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
