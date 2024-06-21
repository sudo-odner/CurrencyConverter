package protHttp

import (
	"encoding/json"
	"fmt"
	"github.com/sudo-odner/CurrencyConverter/internal/entity"
	"io"
	"net/http"
)

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

func (c *HttpClient) GetAllFiat() entity.DataFiat {
	url := c.urlGetFiat
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accepts", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		fmt.Println(err)
		return entity.DataFiat{
			Data: make([]entity.ItemRequest, 0),
		}
	}
	defer resp.Body.Close()

	bodyMap, err := convert(resp.Body)
	if err != nil {
		fmt.Println(err)
		return entity.DataFiat{
			Data: make([]entity.ItemRequest, 0),
		}
	}

	dataMap := bodyMap["data"].([]interface{})

	data := make([]entity.ItemRequest, 0, len(dataMap))
	for _, item := range dataMap {
		var itemStruct entity.ItemRequest
		itemStruct.ID = item.(map[string]any)["id"].(float64)
		itemStruct.Name = item.(map[string]any)["name"].(string)
		itemStruct.Symbol = item.(map[string]any)["symbol"].(string)
		data = append(data, itemStruct)
	}

	return entity.DataFiat{
		Data: data,
	}
}

func (c *HttpClient) GetAllCryptocurrencies() entity.DataCryptocurrencies {
	url := c.urlGetCryptocurrencies
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accepts", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		fmt.Println(err)
		return entity.DataCryptocurrencies{
			Data: make([]entity.ItemRequest, 0),
		}
	}
	defer resp.Body.Close()

	// Декодирование Body в json формат
	bodyMap, err := convert(resp.Body)
	if err != nil {
		fmt.Println(err)
		return entity.DataCryptocurrencies{
			Data: make([]entity.ItemRequest, 0),
		}
	}

	dataMap := bodyMap["data"].([]interface{})

	data := make([]entity.ItemRequest, 0, len(dataMap))
	for _, item := range dataMap {
		var itemStruct entity.ItemRequest
		itemStruct.ID = item.(map[string]any)["id"].(float64)
		itemStruct.Name = item.(map[string]any)["name"].(string)
		itemStruct.Symbol = item.(map[string]any)["symbol"].(string)
		data = append(data, itemStruct)
	}
	return entity.DataCryptocurrencies{
		Data: data,
	}
}

func (c *HttpClient) ConvertOneToOne(amount, from, to float64) entity.ConvertOneToOneRes {
	url := fmt.Sprintf(c.urlConvertOneToOne+"?amount=%.2f&id=%.0f&convert_id=%.0f", amount, from, to)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accepts", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		fmt.Println(err)
		return entity.ConvertOneToOneRes{}
	}
	defer resp.Body.Close()

	// Декодирование Body в json формат
	bodyMap, err := convert(resp.Body)
	if err != nil {
		fmt.Println(err)
		return entity.ConvertOneToOneRes{}
	}

	dataMap := bodyMap["data"].(map[string]interface{})

	strTo := fmt.Sprintf("%.0f", to)
	return entity.ConvertOneToOneRes{
		FromID:     from,
		FromAmount: dataMap["amount"].(float64),
		ToID:       to,
		ToAmount:   dataMap["quote"].(map[string]interface{})[strTo].(map[string]interface{})["price"].(float64),
	}
}
