// +build testnet

package app

import (
	"flag"
	"log"
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	CoinType       = 0
	FundraiserPath = "44'/0'/0'/0/0"
)

var (
	AccountAddressPrefix   = "tcro"
	AccountPubKeyPrefix    = "tcropub"
	ValidatorAddressPrefix = "tcrocncl"
	ValidatorPubKeyPrefix  = "tcrocnclpub"
	ConsNodeAddressPrefix  = "tcrocnclcons"
	ConsNodePubKeyPrefix   = "tcrocnclconspub"
	HumanCoinUnit          = "tcro"
	BaseCoinUnit           = "basetcro" // 10^-8 AKA "carson"
	CroExponent            = 8
	CoinToBaseUnitMuls     = map[string]uint64{
		"tcro": 1_0000_0000,
	}
)

func SetConfig() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(AccountAddressPrefix, AccountPubKeyPrefix)
	config.SetBech32PrefixForValidator(ValidatorAddressPrefix, ValidatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(ConsNodeAddressPrefix, ConsNodePubKeyPrefix)

	config.SetCoinType(CoinType)
	config.SetFullFundraiserPath(FundraiserPath)

	croUnit := sdk.OneDec()
	err := sdk.RegisterDenom(HumanCoinUnit, croUnit)
	if err != nil {
		log.Fatal(err)
	}

	carsonUnit := sdk.NewDecWithPrec(1, int64(CroExponent)) // 10^-8 (carson)
	err = sdk.RegisterDenom(BaseCoinUnit, carsonUnit)

	if err != nil {
		log.Fatal(err)
	}

	config.Seal()
}

var testingConfigState = struct {
	mtx   sync.Mutex
	isSet bool
}{
	isSet: false,
}

func SetTestingConfig() {
	if !isGoTest() {
		panic("SetTestingConfig called but not running go test")
	}

	testingConfigState.mtx.Lock()
	defer testingConfigState.mtx.Unlock()

	if testingConfigState.isSet {
		return
	}

	SetConfig()

	testingConfigState.isSet = true
}

func isGoTest() bool {
	return flag.Lookup("test.v") != nil
}
