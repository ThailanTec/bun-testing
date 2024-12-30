package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/ThailanTec/bun-testing/infra/config"
	"github.com/ThailanTec/bun-testing/infra/db"
	"github.com/ThailanTec/bun-testing/model"
	"github.com/ThailanTec/bun-testing/repository"
	"github.com/google/uuid"
)

func main() {
	runMigrations := flag.Bool("migrate", false, "Run database migrations")
	flag.Parse()

	// Load configuration
	cfg := config.LoadConfig()

	// Initialize Postgres client
	postgresClient, err := db.PostgresClient(cfg)
	if err != nil {
		log.Fatalf("Error initializing Postgres client: %v", err)
	}

	postgresClient.RegisterModel((*model.UserProducts)(nil))

	if *runMigrations {
		db.Migration(postgresClient)
		log.Println("Migrations ran successfully")
		return
	}

	log.Println("Application started without running migrations")

	ctx := context.Background()
	repo := repository.NewQueries(postgresClient)

	repo.CreateProduct(ctx, model.Product{
		Name:  "Samsung Galaxy S21",
		Price: 5000,
	})

	d, err := repo.GetProductsByCategoryID(ctx, uuid.MustParse("530dc81f-1e14-46fa-947c-8b4f144cf9ae"))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	a := fmt.Sprintf("%+v", d)

	fmt.Println(a)
}
