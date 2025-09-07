package model

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type TableName interface {
	TableName() string
}

var DefaultEnum EnumMigrate

type EnumMigrate struct {
	OperationTypes []OperationType `yaml:"operation_types"`
}

func getTbList(enums EnumMigrate) []interface{} {
	var tbList []interface{}
	for _, obj := range enums.OperationTypes {
		tbObj := obj
		tbList = append(tbList, &tbObj)
	}

	return tbList
}

func GetEnumList() ([]interface{}, error) {
	var enums EnumMigrate
	// Get id and values from yaml file
	yamlFile, err := ioutil.ReadFile("enum.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &enums)
	if err != nil {
		return nil, err
	}
	DefaultEnum = enums

	return getTbList(enums), nil
}
