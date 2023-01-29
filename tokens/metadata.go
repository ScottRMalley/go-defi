package tokens

import (
	"strings"

	"github.com/fatih/color"
	cli "github.com/jawher/mow.cli"
	"github.com/scottrmalley/go-evm-constants/networks"
	tl "github.com/scottrmalley/go-token-lists"
	"github.com/scottrmalley/go-token-lists/metadata"

	"github.com/scottrmalley/go-defi/shared"
)

func metadataCommand(cmd *cli.Cmd) {
	networkString := cmd.StringOpt("n network", "", "filter for network")
	showLists := cmd.BoolOpt("l lists", false, "show available token lists")
	list := cmd.StringArg("LIST", "", "list to filter on")
	symbol := cmd.StringArg("SYMBOL", "", "token symbol to fetch")

	cmd.Spec = "[OPTIONS] [LIST] [SYMBOL]"
	cmd.Action = func() {
		if *showLists {
			lists, err := GetAvailableTokenLists(*networkString)
			if err != nil {
				shared.PrintError(err)
				return
			}
			printTokenListMetadata(lists)
			return
		}
		tokens, err := TokensFromList(*networkString, *list, *symbol, false)
		if err != nil {
			shared.PrintError(err)
			return
		}
		if len(tokens) == 0 {
			color.HiYellow("WARN: could not find exact match for symbol %s, attempting partial match\n", *symbol)
			// try again with partial matches
			tokens, err = TokensFromList(*networkString, *list, *symbol, true)
			if err != nil {
				shared.PrintError(err)
				return
			}
		}
		printTokensMetadata(tokens)
	}
}

func filterOnSymbol(symbol string) func(*tl.Token) bool {
	return func(token *tl.Token) bool {
		return strings.ToLower(symbol) == strings.ToLower(token.Symbol)
	}
}

func filterOnChainId(chainId int64) func(*tl.Token) bool {
	return func(token *tl.Token) bool {
		return token.ChainId == chainId
	}
}

func filterOnPartialMatch(symbol string) func(*tl.Token) bool {
	return func(token *tl.Token) bool {
		return strings.Contains(token.Symbol, symbol) || strings.Contains(symbol, token.Symbol)
	}
}

func filterOnPartialMatchOrChainId(chainId int64, symbol string) func(*tl.Token) bool {
	symbolFilter := filterOnPartialMatch(symbol)
	networkFilter := filterOnChainId(chainId)
	return func(token *tl.Token) bool {
		return symbolFilter(token) && networkFilter(token)
	}
}

func filterOnChainIdAndSymbol(chainId int64, symbol string) func(*tl.Token) bool {
	symbolFilter := filterOnSymbol(symbol)
	networkFilter := filterOnChainId(chainId)
	return func(token *tl.Token) bool {
		return symbolFilter(token) && networkFilter(token)
	}
}

func TokensFromList(networkString string, listNameString string, symbol string, allowPartial bool) ([]tl.Token, error) {
	listName, err := metadata.NameFrom(listNameString)
	if err != nil {
		return nil, err
	}
	list, err := tl.NewList(listName)
	if err != nil {
		return nil, err
	}
	if networkString == "" {
		if allowPartial {
			return list.FilterBy(filterOnPartialMatch(symbol)), nil
		}
		return list.FilterBy(filterOnSymbol(symbol)), nil
	}
	network, ok := networks.Supported(networkString)
	if !ok {
		return nil, shared.ErrNetworkNotSupported
	}
	chainId, err := networks.ChainIdFrom(network)
	if err != nil {
		return nil, err
	}
	if allowPartial {
		return list.FilterBy(filterOnPartialMatchOrChainId(chainId.Int64(), symbol)), nil
	}
	return list.FilterBy(filterOnChainIdAndSymbol(chainId.Int64(), symbol)), nil
}

func GetAvailableTokenLists(networkString string) ([]metadata.Name, error) {
	if networkString == "" {
		return metadata.AllLists(), nil
	}
	network, ok := networks.Supported(networkString)
	if !ok {
		return nil, shared.ErrNetworkNotSupported
	}
	return metadata.SupportedLists(network), nil
}
