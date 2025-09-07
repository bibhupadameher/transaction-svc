package postgres

import (
	"context"
	"tx-api/model"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockDBService struct {
	mock.Mock
}

func (m *MockDBService) CreateSchema(schema string) error {
	args := m.Called(schema)
	return args.Error(0)
}

func (m *MockDBService) SetSchema(schema string) error {
	args := m.Called(schema)
	return args.Error(0)
}

func (m *MockDBService) MigrateTables(tables ...interface{}) error {
	args := m.Called(tables)
	return args.Error(0)
}

func (m *MockDBService) MigrateEnums(tables ...interface{}) error {
	args := m.Called(tables)
	return args.Error(0)
}

func (m *MockDBService) BatchWriteData(ctx context.Context, saveddata []model.TableName, deletedData ...model.TableName) error {
	args := m.Called(ctx, saveddata, deletedData)
	return args.Error(0)
}

func (m *MockDBService) FindRows(ctx context.Context, dest interface{}, query func(*gorm.DB) *gorm.DB) error {
	args := m.Called(ctx, dest, query)
	return args.Error(0)
}

func (m *MockDBService) FindFirst(ctx context.Context, dest interface{}, query func(*gorm.DB) *gorm.DB) error {
	args := m.Called(ctx, dest, query)
	return args.Error(0)
}
