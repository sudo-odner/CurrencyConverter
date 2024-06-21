package table

import "github.com/sudo-odner/CurrencyConverter/internal/model"

type ITable interface {
	Add(item model.Item)
	Delete(id float64)
}

type Table struct {
	db []itemHashTable
}

type itemHashTable struct {
	ID                     float64 `json:"id"`
	Name                   string  `json:"name"`
	Symbol                 string  `json:"symbol"`
	CryptocurrenciesOrFiat string  `json:"cryptocurrencies_or_fiat"`
}

func remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}

func (t *Table) Add(item model.Item) {
	newItem := itemHashTable{
		ID:                     item.ID,
		Name:                   item.Name,
		Symbol:                 item.Symbol,
		CryptocurrenciesOrFiat: item.CryptocurrenciesOrFiat,
	}
	t.db = append(t.db, newItem)
}

func (t *Table) Delete(id float64) {
	for idx, item := range t.db {
		if id == item.ID {
			remove(t.db, idx)
		}
	}
}

func New() ITable {
	newHashTable := make([]itemHashTable, 0, 200)

	return &Table{
		db: newHashTable,
	}
}
