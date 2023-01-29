package main

import (
	cli "github.com/jawher/mow.cli"
	"github.com/scottrmalley/go-defi/exchange"
	"github.com/scottrmalley/go-defi/tokens"
	"log"
	"os"
)

func main() {
	app := cli.App("defi", "A tool for querying decentralized finance protocols")
	app.Version("v version", "defi 0.0.1")
	app.Command("token", "Interact with ERC20 tokens", tokens.Command)
	app.Command("dex", "Interact with decentralized exchanges", exchange.Command)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
