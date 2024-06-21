package entity

// Database

type DataCryptocurrenciesAndFiat struct {
	Data []Item `json:"data"`
}

type Item struct {
	ID                     float64 `json:"id"`
	Name                   string  `json:"name"`
	Symbol                 string  `json:"symbol"`
	CryptocurrenciesOrFiat string  `json:"cryptocurrencies_or_fiat"`
}
