package usecase

import "github.com/sudo-odner/CurrencyConverter/internal/entity"

func FiatItemRequestToItem(itemRequest entity.ItemRequest) entity.Item {
	return entity.Item{
		ID:                     itemRequest.ID,
		Name:                   itemRequest.Name,
		Symbol:                 itemRequest.Symbol,
		CryptocurrenciesOrFiat: "fiat",
	}
}

func CryptocurrenciesItemRequestToItem(itemRequest entity.ItemRequest) entity.Item {
	return entity.Item{
		ID:                     itemRequest.ID,
		Name:                   itemRequest.Name,
		Symbol:                 itemRequest.Symbol,
		CryptocurrenciesOrFiat: "cryptocurrencies",
	}
}
