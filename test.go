package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Mega currency converter")
	fmt.Println("For currency write: <amount> <from symbol> <symbol>")
	fmt.Println("For example: 124.4 USDT BTC")
	fmt.Println("---------------------")
	flag := true
	for flag {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		data := strings.Split(text, " ")
		if len(data) == 3 {
			fmt.Println(data[0])
			fmt.Println(data[1])
			fmt.Println(data[2])
		} else {
			fmt.Println("Use example: <amount> <from symbol> <symbol>")
		}

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}

	}

}
