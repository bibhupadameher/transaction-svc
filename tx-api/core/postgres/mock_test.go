package postgres

import (
	"context"
	"errors"
	"testing"
	"tx-api/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestMockDBService_CreateSchema(t *testing.T) {
	mockDB := new(MockDBService)
	mockDB.On("CreateSchema", "testschema").Return(nil)

	err := mockDB.CreateSchema("testschema")
	assert.NoError(t, err)

	mockDB.AssertCalled(t, "CreateSchema", "testschema")
}

func TestMockDBService_SetSchema_Error(t *testing.T) {
	mockDB := new(MockDBService)
	mockDB.On("SetSchema", "schema1").Return(errors.New("set schema failed"))

	err := mockDB.SetSchema("schema1")
	assert.Error(t, err)
	assert.Equal(t, "set schema failed", err.Error())
	mockDB.AssertCalled(t, "SetSchema", "schema1")
}

func TestMockDBService_BatchWriteData(t *testing.T) {
	mockDB := new(MockDBService)
	ctx := context.Background()
	saved := []model.TableName{}
	deleted := []model.TableName{}

	mockDB.On("BatchWriteData", ctx, saved, deleted).Return(nil)

	err := mockDB.BatchWriteData(ctx, saved, deleted...)
	assert.NoError(t, err)
	mockDB.AssertCalled(t, "BatchWriteData", ctx, saved, deleted)
}

func TestMockDBService_FindRows(t *testing.T) {
	mockDB := new(MockDBService)
	dest := []model.TableName{}
	query := func(db *gorm.DB) *gorm.DB { return db }

	mockDB.On("FindRows", context.Background(), dest, mock.AnythingOfType("func(*gorm.DB) *gorm.DB")).Return(nil)

	err := mockDB.FindRows(context.Background(), dest, query)
	assert.NoError(t, err)
	mockDB.AssertCalled(t, "FindRows", context.Background(), dest, mock.AnythingOfType("func(*gorm.DB) *gorm.DB"))
}

func TestMockDBService_FindFirst(t *testing.T) {
	mockDB := new(MockDBService)
	dest := []model.TableName{}
	query := func(db *gorm.DB) *gorm.DB { return db }

	mockDB.On("FindFirst", context.Background(), dest, mock.AnythingOfType("func(*gorm.DB) *gorm.DB")).Return(nil)

	err := mockDB.FindFirst(context.Background(), dest, query)
	assert.NoError(t, err)
	mockDB.AssertCalled(t, "FindFirst", context.Background(), dest, mock.AnythingOfType("func(*gorm.DB) *gorm.DB"))
}
