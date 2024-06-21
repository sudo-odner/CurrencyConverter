package protHttp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DataFiat struct {
	Data []itemRequest `json:"data"`
}

type DataCryptocurrencies struct {
	Data []itemRequest `json:"data"`
}

type itemRequest struct {
	ID     float64 `json:"id"`
	Name   string  `json:"name"`
	Symbol string  `json:"symbol"`
}

// Декодирование Body в map[string]interface{} формат
func convert(body io.ReadCloser) (map[string]interface{}, error) {
	bodyJson, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var bodyMap map[string]interface{}
	if err := json.Unmarshal(bodyJson, &bodyMap); err != nil {
		return nil, err
	}

	return bodyMap, nil
}

func (c Client) GetAllFiat() DataFiat {
	req, _ := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/fiat/map", nil)
	req.Header.Set("Accepts", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		fmt.Println(err)
		return DataFiat{
			Data: make([]itemRequest, 0),
		}
	}
	defer resp.Body.Close()

	bodyMap, err := convert(resp.Body)
	if err != nil {
		fmt.Println(err)
		return DataFiat{
			Data: make([]itemRequest, 0),
		}
	}

	dataMap := bodyMap["data"].([]interface{})

	data := make([]itemRequest, 0, len(dataMap))
	for _, item := range dataMap {
		var itemStruct itemRequest
		itemStruct.ID = item.(map[string]any)["id"].(float64)
		itemStruct.Name = item.(map[string]any)["name"].(string)
		itemStruct.Symbol = item.(map[string]any)["symbol"].(string)
		data = append(data, itemStruct)
	}

	return DataFiat{
		Data: data,
	}
}

func (c Client) GetAllCryptocurrencies() DataCryptocurrencies {
	req, _ := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/map", nil)
	req.Header.Set("Accepts", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		fmt.Println(err)
		return DataCryptocurrencies{
			Data: make([]itemRequest, 0),
		}
	}
	defer resp.Body.Close()

	// Декодирование Body в json формат
	bodyMap, err := convert(resp.Body)
	if err != nil {
		fmt.Println(err)
		return DataCryptocurrencies{
			Data: make([]itemRequest, 0),
		}
	}

	dataMap := bodyMap["data"].([]interface{})

	data := make([]itemRequest, 0, len(dataMap))
	for _, item := range dataMap {
		var itemStruct itemRequest
		itemStruct.ID = item.(map[string]any)["id"].(float64)
		itemStruct.Name = item.(map[string]any)["name"].(string)
		itemStruct.Symbol = item.(map[string]any)["symbol"].(string)
		data = append(data, itemStruct)
	}
	return DataCryptocurrencies{
		Data: data,
	}
}

func (c Client) ConvertOneToOne() {
	//TODO implement me
	panic("implement me")
}
