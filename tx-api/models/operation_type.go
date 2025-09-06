package model

const TableNameOperationType = "operation_type"

type OperationType struct {
	OperationTypeID int    `gorm:"column:operation_type_id;primaryKey" json:"operationTypeID" yaml:"operation_type_id"`
	Description     string `gorm:"column:description;not null" json:"description" yaml:"description"`
}

// TableName OperationType's table name
func (OperationType) TableName() string {
	return TableNameOperationType
}
