package db

import (
	"ToDoAppGo/app/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetInstance(cfg *config.Config) (*gorm.DB, error) {
	if db == nil {
		err := initDB(cfg)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func initDB(cfg *config.Config) error {
	dsn := getDSN(cfg)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func getDSN(cfg *config.Config) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Kyiv",
		cfg.DbHost,
		cfg.DbUser,
		cfg.DbPassword,
		cfg.DbName,
		cfg.DbPort,
	)
}
