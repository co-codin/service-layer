package main

import (
	"context"
	"fmt"
	"log"

	"github.com/co-codin/service-layer/internal/comment"
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

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database.")
		return err
	}

	// cmtService := comment.NewService(db)

	// cmtService.Store.PostComment(
	// 	context.Background()
	// )

	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	fmt.Println("successfully connected and pinged database")

	return nil
}

func main() {
	fmt.Println("rest api")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
