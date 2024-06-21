package protHttp

import (
	"github.com/sudo-odner/CurrencyConverter/internal/controler/protHttp/middleware"
	"github.com/sudo-odner/CurrencyConverter/internal/entity"
	"net/http"
)

type IHttpClient interface {
	GetAllCryptocurrencies() entity.DataCryptocurrencies
	GetAllFiat() entity.DataFiat
	ConvertOneToOne(amount, from, to float64) entity.ConvertOneToOneRes
}

type HttpClient struct {
	client                 http.Client
	urlGetFiat             string
	urlGetCryptocurrencies string
	urlConvertOneToOne     string
}

func New(urlGetFiat, urlGetCryptocurrencies, urlConvertOneToOne string) IHttpClient {
	stackMiddleware := middleware.CreateStack(
		middleware.SecretKey,
		middleware.ResponseStatus,
	)
	newHttpClient := http.Client{
		Transport: stackMiddleware(http.DefaultTransport),
	}

	return &HttpClient{
		client:                 newHttpClient,
		urlGetFiat:             urlGetFiat,
		urlGetCryptocurrencies: urlGetCryptocurrencies,
		urlConvertOneToOne:     urlConvertOneToOne,
	}
}
