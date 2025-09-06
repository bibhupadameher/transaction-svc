package postgres

import (
	"context"
	"fmt"
	"os"
	"tx-api/config"
	model "tx-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresService struct {
	DB *gorm.DB
}

func NewPostgresService() (*PostgresService, error) {
	gormConfig := &gorm.Config{}

	env := os.Getenv("APP_ENV")
	if env != "prod" {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	conf, err := config.Get()
	if err != nil {
		return nil, fmt.Errorf("failed to get config: %w", err)
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.Database.Host, conf.Database.Port, conf.Database.User, conf.Database.Password, conf.Database.Name,
	)

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect db: %w", err)
	}

	service := &PostgresService{DB: db}

	if err := service.CreateSchema(conf.Database.Schema); err != nil {
		return nil, fmt.Errorf("failed to create schema: %w", err)
	}

	if err := service.SetSchema(conf.Database.Schema); err != nil {
		return nil, fmt.Errorf("failed to set schema: %w", err)
	}

	tables := model.GetModelList()
	if err := service.MigrateTables(tables...); err != nil {
		return nil, fmt.Errorf("failed to migrate tables: %w", err)
	}

	enumsList, err := model.GetEnumList()
	if err != nil {
		return nil, fmt.Errorf("failed to get enum list : %w", err)
	}
	if err := service.MigrateEnums(enumsList...); err != nil {
		return nil, fmt.Errorf("failed to migrate enum: %w", err)
	}

	return service, nil
}

func (p *PostgresService) CreateSchema(schema string) error {
	return p.DB.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schema)).Error
}

func (p *PostgresService) SetSchema(schema string) error {
	return p.DB.Exec(fmt.Sprintf("SET search_path TO %s", schema)).Error
}

func (p *PostgresService) MigrateTables(tables ...interface{}) error {
	return p.DB.AutoMigrate(tables...)
}

func (p *PostgresService) BatchWriteData(ctx context.Context, data ...interface{}) error {
	return p.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, d := range data {
			if err := tx.Create(d).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (p *PostgresService) FindRows(ctx context.Context, dest interface{}, query func(*gorm.DB) *gorm.DB) error {
	return query(p.DB.WithContext(ctx)).Find(dest).Error
}

func (p *PostgresService) MigrateEnums(tables ...interface{}) error {
	if len(tables) == 0 {
		return nil
	}

	return p.DB.Transaction(func(tx *gorm.DB) error {
		for _, tbl := range tables {

			if err := tx.Model(tbl).Create(tbl).Error; err != nil {
				return err
			}

		}
		return nil
	})
}
