package table

import "github.com/sudo-odner/CurrencyConverter/internal/entity"

type ITable interface {
	Add(item entity.Item)
	Find(symbol string) *entity.Item
	Delete(id float64)
}

type Table struct {
	db []entity.Item
}

func New() ITable {
	newHashTable := make([]entity.Item, 0, 200)

	return &Table{
		db: newHashTable,
	}
}
