package adapters

import (
	"fmt"
	"rices/common/configs"
	"rices/core/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Pgsql struct {
	db *gorm.DB
}

func NewPgsql() *Pgsql {
	return &Pgsql{}
}

func (p *Pgsql) Connect() error {
	dataSource := configs.Get().DataSource

	db, err := gorm.Open(postgres.Open(dataSource), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	db.AutoMigrate(&domain.Users{})
	db.AutoMigrate(&domain.FileDescriptors{})
	p.db = db
	return nil
}

func (p *Pgsql) DB() *gorm.DB {
	return p.db
}
