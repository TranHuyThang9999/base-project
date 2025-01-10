package adapters

import (
	"fmt"
	"rices/common/configs"
	"rices/core/domain"
	"time"

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

	// Add connection timeout and other parameters
	dataSourceWithParams := fmt.Sprintf("%s connect_timeout=10 statement_timeout=10000 idle_in_transaction_session_timeout=10000",
		dataSource)

	config := &gorm.Config{
		PrepareStmt: true,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	db, err := gorm.Open(postgres.Open(dataSourceWithParams), config)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %v", err)
	}

	// Configure connection pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	// Auto migrate schemas
	if err := db.AutoMigrate(&domain.Users{},
		&domain.FileDescriptors{}); err != nil {
		return fmt.Errorf("failed to auto migrate schemas: %v", err)
	}

	p.db = db
	return nil
}

func (p *Pgsql) DB() *gorm.DB {
	return p.db
}
