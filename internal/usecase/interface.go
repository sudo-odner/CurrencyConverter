package usecase

import (
	"github.com/sudo-odner/CurrencyConverter/internal/entity"
)

type (
	// IHttpClient - http client
	IHttpClient interface {
		GetAllCryptocurrencies() entity.DataCryptocurrencies
		GetAllFiat() entity.DataFiat
		ConvertOneToOne(amount, from, to float64) entity.ConvertOneToOneRes
	}

	// ITable - DB table
	ITable interface {
		Add(item entity.Item)
		Find(symbol string) *entity.Item
		Delete(id float64)
	}
)
