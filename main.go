package main

import (
	"fmt"

	"github.com/charoleizer/tadashi-bot/bot"
	"github.com/charoleizer/tadashi-bot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
