package main

import "github.com/sudo-odner/CurrencyConverter/internal/controler/protHttp"

func main() {
	newCLient := protHttp.New()
	newCLient.GetAllCoin()
}
