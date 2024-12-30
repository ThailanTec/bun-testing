package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Users struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID        uuid.UUID  `bun:"type:uuid,default:gen_random_uuid(),pk"`
	Name      string     `bun:"name,notnull"`
	CreatedAt time.Time  `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time  `bun:",nullzero,notnull,default:current_timestamp"`
	Products  []*Product `bun:"m2m:user_products,join:User=Product"`
}

type UserProducts struct {
	bun.BaseModel `bun:"table:user_products,alias:up"`

	UserID    uuid.UUID `bun:"type:uuid,notnull"`
	ProductID uuid.UUID `bun:"type:uuid,notnull"`
	User      *Users    `bun:"rel:belongs-to,join:user_id=id"`
	Product   *Product  `bun:"rel:belongs-to,join:product_id=id"`
}
