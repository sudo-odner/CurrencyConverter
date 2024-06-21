package table

import "github.com/sudo-odner/CurrencyConverter/internal/entity"

func remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}

func (t *Table) Add(item entity.Item) {
	newItem := entity.Item{
		ID:                     item.ID,
		Name:                   item.Name,
		Symbol:                 item.Symbol,
		CryptocurrenciesOrFiat: item.CryptocurrenciesOrFiat,
	}
	t.db = append(t.db, newItem)
}

func (t *Table) Find(symbol string) *entity.Item {
	for _, item := range t.db {
		if symbol == item.Symbol {
			return &item
		}
	}
	return nil
}

func (t *Table) Delete(id float64) {
	for idx, item := range t.db {
		if id == item.ID {
			remove(t.db, idx)
		}
	}
}
