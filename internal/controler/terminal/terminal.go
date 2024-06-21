package terminal

import (
	"bufio"
	"fmt"
	"github.com/sudo-odner/CurrencyConverter/internal/usecase"
	"os"
	"strconv"
	"strings"
)

type Terminal struct {
	us usecase.UseCase
}

func New(us usecase.UseCase) Terminal {
	return Terminal{
		us: us,
	}
}

func (t *Terminal) Start() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Mega currency converter")
	fmt.Println("For currency write: <amount> <from symbol> <symbol>")
	fmt.Println("For example: 124.4 USDT BTC")
	fmt.Println("Use 'quit' or 'q' for exit")
	fmt.Println("---------------------")
	flag := true
	for flag {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')

		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		if strings.Compare("quit", text) == 0 || strings.Compare("q", text) == 0 {
			flag = false
			fmt.Println("---------------------")
			fmt.Println("Exit")
		}
		data := strings.Split(text, " ")
		if len(data) == 3 {
			fl, err := strconv.ParseFloat(data[0], 64)
			if err != nil {
				fmt.Println("amount - is not float")
			} else {
				fmt.Println(t.us.ConvertOneToOne(fl, data[1], data[2]))
			}
		} else {
			fmt.Println("Use example: <amount> <from symbol> <symbol>")
		}
	}
}
