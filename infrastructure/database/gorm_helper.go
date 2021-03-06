package database

import (
	"context"
	"gorm.io/gorm"
	"mywallet/application/apperror"
)

type contextDBType string

var contextDBValue contextDBType = "DB"

// ExtractDB is used by other repo to extract the database from context
func ExtractDB(ctx context.Context) (*gorm.DB, error) {

	db, ok := ctx.Value(contextDBValue).(*gorm.DB)
	if !ok {
		return nil, apperror.DatabaseNotFoundInContextError
	}

	return db, nil
}
