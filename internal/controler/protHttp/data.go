package protHttp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DataCryptocurrenciesAndFiat struct {
	Data []itemRequest `json:"data"`
}
type itemRequest struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

func (c Client) GetAllCoin() {
	req, _ := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/fiat/map?saf", nil)
	req.Header.Set("Accepts", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		return
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)

	// Декодирование Body в json формат
	bodyJson, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	// Закрытие Body
	defer resp.Body.Close()

	var x map[string]interface{}

	json.Unmarshal(bodyJson, &x)
	//fmt.Println(x)
	for _, item := range x["data"].([]interface{}) {
		fmt.Println(item)
	}

}

func (c Client) GetAllFiat() {
	//TODO implement me
	panic("implement me")
}

func (c Client) ConvertOneToOne() {
	//TODO implement me
	panic("implement me")
}
