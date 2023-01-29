package tokens

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/scottrmalley/go-defi/shared"
	"github.com/scottrmalley/go-evm-constants/networks"
	tl "github.com/scottrmalley/go-token-lists"
	"github.com/scottrmalley/go-token-lists/metadata"
	"math/big"
	"strings"
)

func printTokenListMetadata(names []metadata.Name) {
	for _, name := range names {
		fmt.Print("\n")
		data, err := metadata.TokenListMetadata(name)
		if err != nil {
			shared.PrintError(err)
			continue
		}
		color.HiGreen("===== %s =====\n", strings.ToUpper(string(data.Name)))
		color.Cyan("name:\t%s\n", data.Name)
		color.Cyan("supported networks:")
		for _, net := range data.Networks {
			color.Magenta("\t\t%s\n", net)
		}
		color.Cyan("url: \t%s\n", data.Url)
	}
	fmt.Print("\n")
}

func printTokensMetadata(tokens []tl.Token) {
	if len(tokens) < 1 {
		color.HiYellow("WARN: no tokens found for given filters")
	}
	for _, token := range tokens {
		network, err := networks.NameFrom(big.NewInt(token.ChainId))
		if err != nil {
			network = "UNKNOWN"
		}
		fmt.Print("\n")
		color.HiGreen("===== %s =====\n", strings.ToUpper(token.Symbol))
		color.Cyan("name:\t%s\n", token.Name)
		color.Cyan("symbol:\t%s\n", token.Symbol)
		color.Cyan("network:\t%s\n", network)
		color.Cyan("chain id:\t%d\n", token.ChainId)
		color.Cyan("decimals:\t%d\n", token.Decimals)
		color.Cyan("address:\t%s\n", token.Address.Hex())
	}
	fmt.Print("\n")
}
