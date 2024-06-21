package app

import (
	"fmt"
	"github.com/sudo-odner/CurrencyConverter/config"
	"github.com/sudo-odner/CurrencyConverter/internal/controler/protHttp"
	"github.com/sudo-odner/CurrencyConverter/internal/controler/terminal"
	"github.com/sudo-odner/CurrencyConverter/internal/repository/table"
	usecase2 "github.com/sudo-odner/CurrencyConverter/internal/usecase"
)

func Start() {
	// TODO init config
	cfg, err := config.Load()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// TODO init client
	client := protHttp.New(
		cfg.ProtHttp.Url.Url_get_fiat,
		cfg.ProtHttp.Url.Url_get_cryptocurrencies,
		cfg.ProtHttp.Url.Url_convert_one_to_one,
	)
	// TODO init db
	db := table.New()
	// TODO init usecase
	uc := usecase2.New(db, client)
	uc.WriteMapCryptocurrenciesAndFiat()
	// TODO init terminal
	term := terminal.New(*uc)
	term.Start()
}
