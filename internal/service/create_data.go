package service

import (
	"context"

	"github.com/ThailanTec/bun-testing/repository"
	"github.com/google/uuid"
)

type CreateDataInterface interface {
	CreateData(ctx context.Context) error
	RecoveryData() any
}

type createData struct {
	repo repository.Queries
}

func NewCreateData(repo repository.Queries) createData {
	return createData{repo: repo}
}

func (c *createData) CreateData(ctx context.Context) error {

	// Add Product to User
	err := c.repo.AddProductToUser(context.Background(), uuid.MustParse("81559088-3e45-4d0a-b031-896adbbce580"), uuid.MustParse("87141f8b-74fb-4b9a-90ed-49de5ba9d216"))
	if err != nil {
		return err
	}

	return nil
}
