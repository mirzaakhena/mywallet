package database

import (
	"context"
	"gorm.io/gorm"
)

type GormWithoutTrxImpl struct {
	DB *gorm.DB
}

func NewGormWithoutTrxImpl(db *gorm.DB) *GormWithoutTrxImpl {
	return &GormWithoutTrxImpl{DB: db}
}

func (r *GormWithoutTrxImpl) GetDatabase(ctx context.Context) (context.Context, error) {
	trxCtx := context.WithValue(ctx, ContextDBValue, r.DB)
	return trxCtx, nil
}
