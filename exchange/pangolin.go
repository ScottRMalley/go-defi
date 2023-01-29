package exchange

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/scottrmalley/go-evm-constants/contracts/dex"
	"github.com/scottrmalley/go-evm-constants/networks"
	tl "github.com/scottrmalley/go-token-lists"
	"github.com/scottrmalley/go-token-lists/metadata"
	"github.com/shopspring/decimal"

	"github.com/scottrmalley/go-defi/bindings"
	"github.com/scottrmalley/go-defi/shared"
)

type PangolinClient struct {
	client    *ethclient.Client
	network   networks.Name
	factory   *bindings.IPangolinFactory
	tokenList *tl.List
}

func NewPangolinClient(network networks.Name) (*PangolinClient, error) {
	networkMetadata, err := networks.GetNetwork(network)
	if err != nil {
		return nil, err
	}

	factoryAddress, err := dex.GetContracts(dex.PANGOLIN, network)
	if err != nil {
		return nil, err
	}

	client, err := ethclient.Dial(networkMetadata.RpcUrl)
	if err != nil {
		return nil, err
	}

	factory, err := bindings.NewIPangolinFactory(factoryAddress.Factory, client)
	if err != nil {
		return nil, err
	}

	tokenList, err := tl.NewList(metadata.PANGOLIN)
	if err != nil {
		return nil, err
	}

	return &PangolinClient{factory: factory, tokenList: tokenList, network: network, client: client}, nil
}

func (c *PangolinClient) GetPair(baseSymbol string, quoteSymbol string) (common.Address, error) {
	baseToken, quoteToken, err := c.getTokens(baseSymbol, quoteSymbol)
	if err != nil {
		return common.Address{}, nil
	}
	return c.factory.GetPair(nil, baseToken.Address, quoteToken.Address)
}

func (c *PangolinClient) GetPrice(baseSymbol string, quoteSymbol string) (decimal.Decimal, error) {
	baseToken, quoteToken, err := c.getTokens(baseSymbol, quoteSymbol)
	if err != nil {
		return decimal.Zero, err
	}
	baseReserve, quoteReserve, err := c.GetReserves(baseSymbol, quoteSymbol)
	if err != nil {
		return decimal.Zero, err
	}
	return decimal.NewFromBigInt(quoteReserve, -int32(quoteToken.Decimals)).
		Div(decimal.NewFromBigInt(baseReserve, -int32(baseToken.Decimals))), nil
}

func (c *PangolinClient) GetReserves(baseSymbol string, quoteSymbol string) (*big.Int, *big.Int, error) {
	baseToken, _, err := c.getTokens(baseSymbol, quoteSymbol)
	if err != nil {
		return nil, nil, err
	}

	pairAddress, err := c.GetPair(baseSymbol, quoteSymbol)
	if err != nil {
		return nil, nil, err
	}

	pair, err := bindings.NewIPangolinPair(pairAddress, c.client)
	if err != nil {
		return nil, nil, err
	}

	token0, err := pair.Token0(nil)
	if err != nil {
		return nil, nil, err
	}

	reserves, err := pair.GetReserves(nil)
	if err != nil {
		return nil, nil, err
	}

	// check pair ordering
	var baseReserve, quoteReserve *big.Int
	if token0 == baseToken.Address {
		baseReserve = reserves.Reserve0
		quoteReserve = reserves.Reserve1
	} else {
		baseReserve = reserves.Reserve1
		quoteReserve = reserves.Reserve0
	}
	return baseReserve, quoteReserve, nil
}

func (c *PangolinClient) getTokens(baseSymbol string, quoteSymbol string) (*tl.Token, *tl.Token, error) {
	baseToken, err := c.tokenList.TokenBySymbol(c.network, baseSymbol)
	if err != nil {
		return nil, nil, err
	}
	quoteToken, err := c.tokenList.TokenBySymbol(c.network, quoteSymbol)
	if err != nil {
		return nil, nil, err
	}
	return baseToken, quoteToken, nil
}

func (c *PangolinClient) FormatPrice(price *big.Int, quoteSymbol string) (decimal.Decimal, error) {
	quoteToken, err := c.tokenList.TokenBySymbol(c.network, quoteSymbol)
	if err != nil {
		return decimal.Zero, err
	}
	return shared.ToDecimal(price, quoteToken.Decimals), nil
}
