package model

func GetModelList() []interface{} {
	return []interface{}{
		Transaction{},
		OperationType{},
		Account{},
	}
}
