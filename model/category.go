package model

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Category struct {
	bun.BaseModel `bun:"table:categories,alias:c"`

	ID   uuid.UUID `bun:"type:uuid,default:gen_random_uuid(),pk"`
	Name string    `bun:"name,notnull"`
}
