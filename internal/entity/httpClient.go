package entity

type DataFiat struct {
	Data []ItemRequest `json:"data"`
}

type DataCryptocurrencies struct {
	Data []ItemRequest `json:"data"`
}

type ItemRequest struct {
	ID     float64 `json:"id"`
	Name   string  `json:"name"`
	Symbol string  `json:"symbol"`
}

type ConvertOneToOneRes struct {
	FromID     float64 `json:"from_id"`
	FromAmount float64 `json:"from_amount"`
	ToID       float64 `json:"to_id"`
	ToAmount   float64 `json:"to_amount"`
}
