package db

import (
	"context"
	"log"

	"github.com/ThailanTec/bun-testing/model"
	"github.com/uptrace/bun"
)

func Migration(db *bun.DB) {
	ctx := context.Background()

	// Create Users table
	if _, err := db.NewCreateTable().Model((*model.Users)(nil)).Exec(ctx); err != nil {
		log.Fatalf("Error creating Users table: %v", err)
	}

	// Create Category table
	if _, err := db.NewCreateTable().Model((*model.Category)(nil)).Exec(ctx); err != nil {
		log.Fatalf("Error creating Category table: %v", err)
	}

	// Create Product table with foreign key to Category
	if _, err := db.NewCreateTable().Model((*model.Product)(nil)).
		ForeignKey(`("category_id") REFERENCES "categories" ("id") ON DELETE CASCADE`).
		Exec(ctx); err != nil {
		log.Fatalf("Error creating Product table: %v", err)
	}

	// Create UserProducts table with foreign keys to Users and Products
	if _, err := db.NewCreateTable().Model((*model.UserProducts)(nil)).
		ForeignKey(`("user_id") REFERENCES "users" ("id") ON DELETE CASCADE`).
		ForeignKey(`("product_id") REFERENCES "products" ("id") ON DELETE CASCADE`).
		Exec(ctx); err != nil {
		log.Fatalf("Error creating UserProducts table: %v", err)
	}
}
