package main

import (
	"fmt"
	"github.com/sudo-odner/CurrencyConverter/internal/controler/protHttp"
)

func main() {
	newClient := protHttp.New()
	fmt.Println(newClient.ConvertOneToOne(31, 1, 2781))
}
