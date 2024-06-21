package main

import "github.com/sudo-odner/CurrencyConverter/internal/controler/protHttp"

func main() {
	newClient := protHttp.New()
	newClient.GetAllCryptocurrencies()
}
