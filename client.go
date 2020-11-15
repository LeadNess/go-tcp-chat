package main

import (
	"log"
	"os"

	"github.com/vnkrtv/go-tcp-chat/tui"
)

func main()  {
	client := tui.LoginWindowUI()
	if client == nil {
		os.Exit(1)
	}
	ui := tui.ChatWindowUI(client)
	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}