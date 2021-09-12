package database

import (
  "context"
  "gorm.io/gorm"
  "mywallet/infrastructure/log"
)

type GormWithTrxImpl struct {
  DB *gorm.DB
}

func NewGormWithTrxImpl(db *gorm.DB) *GormWithTrxImpl {
  return &GormWithTrxImpl{DB: db}
}

func (r *GormWithTrxImpl) BeginTransaction(ctx context.Context) (context.Context, error) {

  dbTrx := r.DB.Begin()

  trxCtx := context.WithValue(ctx, ContextDBValue, dbTrx)

  return trxCtx, nil
}

func (r *GormWithTrxImpl) CommitTransaction(ctx context.Context) error {
  log.Info(ctx, "Commit")

  db, err := ExtractDB(ctx)
  if err != nil {
    return err
  }

  return db.Commit().Error
}

func (r *GormWithTrxImpl) RollbackTransaction(ctx context.Context) error {
  log.Info(ctx, "Rollback")

  db, err := ExtractDB(ctx)
  if err != nil {
    return err
  }

  return db.Rollback().Error
}
