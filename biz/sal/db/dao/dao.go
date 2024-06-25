package dao

import (
	"context"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var (
	InitContext = context.Background()
)

var ProvideDao = wire.NewSet(
	New,
)

func New(db *gorm.DB) *Query {
	return Use(db)
}
