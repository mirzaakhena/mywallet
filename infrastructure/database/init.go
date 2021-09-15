package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewGormDefault ...
func NewGormDefault() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("wallet.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
