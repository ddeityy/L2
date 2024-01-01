package api

import (
	"github.com/glebarez/sqlite"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

func ConnectDB(logger *zap.Logger) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{Logger: zapgorm2.New(logger)})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Event{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
