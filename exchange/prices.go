package exchange

import (
	"errors"

	cli "github.com/jawher/mow.cli"
	"github.com/scottrmalley/go-evm-constants/contracts/dex"
	"github.com/scottrmalley/go-evm-constants/networks"
	"github.com/shopspring/decimal"

	"github.com/scottrmalley/go-defi/shared"
)

type PriceFetcher interface {
	GetPrice(string, string) (decimal.Decimal, error)
}

func pricesCommand(cmd *cli.Cmd) {
	networkString := cmd.StringOpt("n network", "", "filter for network")
	dexNameString := cmd.StringArg("DEX", "", "exchange to fetch price from")
	baseSymbol := cmd.StringArg("BASE", "", "base currency symbol")
	quoteSymbol := cmd.StringArg("QUOTE", "USDC", "quote currency symbol")

	cmd.Spec = "[OPTIONS] [DEX] [BASE] [QUOTE]"
	cmd.Action = func() {
		fetcher, err := GetPriceFetcher(dex.Name(*dexNameString), networks.Name(*networkString))
		if err != nil {
			shared.PrintError(err)
			return
		}
		price, err := fetcher.GetPrice(*baseSymbol, *quoteSymbol)
		if err != nil {
			shared.PrintError(err)
			return
		}
		printPriceData(price, *dexNameString, *baseSymbol, *quoteSymbol)
	}
}

func GetPriceFetcher(dexName dex.Name, network networks.Name) (PriceFetcher, error) {
	switch dexName {
	case dex.PANGOLIN:
		return NewPangolinClient(network)
	default:
		return nil, errors.New("no dex found for network or name")
	}
}
