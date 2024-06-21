package usecase

import "fmt"

type UseCase struct {
	iTable     ITable
	httpClient IHttpClient
}

// New -.
func New(r ITable, w IHttpClient) *UseCase {
	return &UseCase{
		iTable:     r,
		httpClient: w,
	}
}

func (u *UseCase) WriteMapCryptocurrenciesAndFiat() {
	//TODO: get map cryptocurrencies and fiat
	dataFiat := u.httpClient.GetAllFiat()
	dataCryptocurrencies := u.httpClient.GetAllCryptocurrencies()
	//TODO: write id db
	for _, item := range dataFiat.Data {
		newItem := FiatItemRequestToItem(item)
		u.iTable.Add(newItem)
	}
	for _, item := range dataCryptocurrencies.Data {
		newItem := CryptocurrenciesItemRequestToItem(item)
		u.iTable.Add(newItem)
	}
}

func (u *UseCase) ConvertOneToOne(amount float64, from, to string) string {
	fmt.Println("es", amount, from, to)
	// TODO find form and to
	formInfo := u.iTable.Find(from)
	toInfo := u.iTable.Find(to)
	if formInfo == nil {
		return "from - does not exist"
	}
	if toInfo == nil {
		return "to - does not exist"
	}
	// TODO get edit
	data := u.httpClient.ConvertOneToOne(amount, formInfo.ID, toInfo.ID)
	// TODO return result
	msgLeft := fmt.Sprintf("%f %s", data.FromAmount, formInfo.Symbol)
	msgRight := fmt.Sprintf("%f %s", data.ToAmount, toInfo.Symbol)
	return msgLeft + " -> " + msgRight
}
