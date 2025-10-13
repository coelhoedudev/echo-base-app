package database

import (
	"fmt"
	"infra-base-go/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB     *gorm.DB
	config *config.DBConfig
}

func New(db *gorm.DB, config *config.DBConfig) *Database {
	return &Database{
		DB:     db,
		config: config,
	}
}

func (b *Database) Connect() (db *gorm.DB, err error) {
	dsn := b.getDsn()
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return
}

func (b *Database) getDsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		b.config.Host,
		b.config.User,
		b.config.Password,
		b.config.Name,
		b.config.Port,
	)
}
