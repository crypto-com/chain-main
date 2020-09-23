package test

import (
	fmt "fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// FakeCoinParser maintains string to sdk.Coin mapping for testing
type FakeCoinParser struct {
	mapping map[string]sdk.Coin
}

func NewFakeCoinParser(mapping map[string]sdk.Coin) *FakeCoinParser {
	return &FakeCoinParser{
		mapping,
	}
}

func (parser *FakeCoinParser) ParseCoin(s string) (sdk.Coin, error) {
	coin, exist := parser.mapping[s]
	if exist {
		return coin, nil
	}
	return sdk.Coin{}, fmt.Errorf("coin mapping not found: %s", s)
}

func (parser *FakeCoinParser) GetBaseUnit() string {
	return "baseunit"
}

// FIXME: NOT IMPLEMENTED
func (parser *FakeCoinParser) MustSprintBaseCoin(baseCoin sdk.Int, denom string) string {
	return ""
}

// FIXME: NOT IMPLEMENTED
func (parser *FakeCoinParser) SprintBaseCoin(baseCoin sdk.Int, denom string) (string, error) {
	return "", nil
}

func (parser *FakeCoinParser) ParseCoinToa(s string) (string, error) {
	coin, err := parser.ParseCoin(s)
	if err != nil {
		return "", err
	}
	return coin.String(), nil
}

func (parser *FakeCoinParser) ParseCoins(s string) (sdk.Coins, error) {
	coinStrs := strings.Split(s, ",")
	coins := make(sdk.Coins, len(coinStrs))

	for i, coinStr := range coinStrs {
		coin, err := parser.ParseCoin(coinStr)
		if err != nil {
			return nil, err
		}

		coins[i] = coin
	}

	// sort coins for determinism
	coins.Sort()

	return coins, nil
}

func (parser *FakeCoinParser) ParseCoinsToa(s string) (string, error) {
	coins, err := parser.ParseCoins(s)
	if err != nil {
		return "", err
	}
	return coins.String(), nil
}
