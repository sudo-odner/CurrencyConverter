package protHttp

import (
	"github.com/sudo-odner/CurrencyConverter/internal/controler/protHttp/middleware"
	"net/http"
)

type IClient interface {
	GetAllCryptocurrencies() DataCryptocurrencies
	GetAllFiat() DataFiat
	ConvertOneToOne()
}

type Client struct {
	client http.Client
}

func New() IClient {
	stackMiddleware := middleware.CreateStack(
		middleware.SecretKey,
		middleware.ResponseStatus,
	)
	newHttpClient := http.Client{
		Transport: stackMiddleware(http.DefaultTransport),
	}

	return &Client{
		client: newHttpClient,
	}
}
