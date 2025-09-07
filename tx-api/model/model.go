package model

func GetModelList() []interface{} {
	return []interface{}{
		OperationType{},
		Account{},
		Transaction{},
	}
}
