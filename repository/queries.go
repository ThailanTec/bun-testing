package repository

import (
	"context"

	"github.com/ThailanTec/bun-testing/model"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Queries interface {
	CreateUser(ctx context.Context, user model.Users) error
	GetUserByID(id uuid.UUID) (model.Users, error)

	CreateProduct(ctx context.Context, prd model.Product) error
	GetProductByID(id uuid.UUID) (model.Product, error)

	CreateCategory(ctx context.Context, category model.Category) error
	GetCategoryByID(id uuid.UUID) (model.Category, error)

	AddProductToUser(ctx context.Context, userID, productID uuid.UUID) error
	GetUserWithProducts(ctx context.Context, userID uuid.UUID) (model.Users, error)

	GetProductsByCategoryID(ctx context.Context, categoryID uuid.UUID) ([]model.Product, error)
}

type queries struct {
	db *bun.DB
}

func NewQueries(db *bun.DB) queries {
	return queries{db: db}
}

func (d *queries) CreateUser(ctx context.Context, user model.Users) error {
	_, err := d.db.NewInsert().Model(&user).Exec(ctx)
	return err
}

func (d *queries) GetUserByID(id uuid.UUID) (model.Users, error) {
	var user model.Users
	err := d.db.NewSelect().Model(&user).Where("id = ?", id).Scan(context.Background())
	return user, err
}

func (d *queries) CreateProduct(ctx context.Context, prd model.Product) error {
	_, err := d.db.NewInsert().Model(&prd).Exec(ctx)
	return err
}

func (d *queries) GetProductByID(id uuid.UUID) (model.Product, error) {
	var prd model.Product
	err := d.db.NewSelect().Model(&prd).Where("id = ?", id).Scan(context.Background())
	return prd, err
}

func (d *queries) CreateCategory(ctx context.Context, category model.Category) error {
	_, err := d.db.NewInsert().Model(&category).Exec(ctx)
	return err
}

func (d *queries) GetCategoryByID(id uuid.UUID) (model.Category, error) {
	var category model.Category
	err := d.db.NewSelect().Model(&category).Where("id = ?", id).Scan(context.Background())
	return category, err
}

func (d *queries) AddProductToUser(ctx context.Context, userID, productID uuid.UUID) error {
	userProduct := model.UserProducts{
		UserID:    userID,
		ProductID: productID,
	}

	_, err := d.db.NewInsert().Model(&userProduct).Exec(ctx)
	return err
}

func (d *queries) GetUserWithProducts(ctx context.Context, userID uuid.UUID) (model.Users, error) {
	var user model.Users
	err := d.db.NewSelect().
		Model(&user).
		Relation("Products").
		Where("u.id = ?", userID).
		Scan(ctx)
	return user, err
}

func (d *queries) GetProductsByCategoryID(ctx context.Context, categoryID uuid.UUID) ([]model.Product, error) {
	var products []model.Product
	err := d.db.NewSelect().
		Model(&products).
		Where("category_id = ?", categoryID).
		Scan(ctx)
	return products, err
}
