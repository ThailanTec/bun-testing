package model

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Product struct {
	bun.BaseModel `bun:"table:products,alias:p"`

	ID         uuid.UUID `bun:"type:uuid,default:gen_random_uuid(),pk"`
	Name       string    `bun:"name,notnull"`
	Price      float64   `bun:"price,notnull"`
	CategoryID uuid.UUID `bun:"type:uuid,notnull"`
	Category   *Category `bun:"rel:belongs-to,join:category_id=id"`
}
