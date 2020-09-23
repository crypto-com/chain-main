package app_test

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	keys "github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/crypto-com/chain-main/app"
)

func TestMnemonic(t *testing.T) {
	kb := keys.NewInMemory()
	account, err := kb.NewAccount(
		"croTest",
		//nolint:lll
		"point shiver hurt flight fun online hub antenna engine pave chef fantasy front interest poem accident catch load frequent praise elite pet remove used",
		"",
		app.FundraiserPath,
		hd.Secp256k1,
	)
	require.NoError(t, err)

	publicKey := account.GetPubKey()
	expectedPublicKey := []byte("0396bb69cbbf27c07e08c0a9d8ac2002ed75a6287a3f2e4cfe11977817ca14fad0")

	expectedPublicKeyBytes := make([]byte, hex.DecodedLen(len(expectedPublicKey)))
	_, err = hex.Decode(expectedPublicKeyBytes, expectedPublicKey)
	require.NoError(t, err)

	if !bytes.Equal(expectedPublicKeyBytes, publicKey.Bytes()) {
		t.Error("HD public key does not match to expected public key")
	}
}

func TestConversion(t *testing.T) {
	app.SetTestingConfig()

	testCases := []struct {
		input  sdk.Coin
		denom  string
		result sdk.Coin
		expErr bool
	}{
		{sdk.NewCoin("foo", sdk.ZeroInt()), app.HumanCoinUnit, sdk.Coin{}, true},
		{sdk.NewCoin(app.HumanCoinUnit, sdk.ZeroInt()), "foo", sdk.Coin{}, true},
		{sdk.NewCoin(app.HumanCoinUnit, sdk.ZeroInt()), "FOO", sdk.Coin{}, true},

		{sdk.NewCoin(app.HumanCoinUnit, sdk.NewInt(5)),
			app.BaseCoinUnit, sdk.NewCoin(app.BaseCoinUnit, sdk.NewInt(500000000)), false}, // cro => carson

		{sdk.NewCoin(app.BaseCoinUnit, sdk.NewInt(500000000)),
			app.HumanCoinUnit, sdk.NewCoin(app.HumanCoinUnit, sdk.NewInt(5)), false}, // carson => cro

	}

	for i, tc := range testCases {
		res, err := sdk.ConvertCoin(tc.input, tc.denom)
		require.Equal(
			t, tc.expErr, err != nil,
			"unexpected error; tc: #%d, input: %s, denom: %s", i+1, tc.input, tc.denom,
		)
		require.Equal(
			t, tc.result, res,
			"invalid result; tc: #%d, input: %s, denom: %s", i+1, tc.input, tc.denom,
		)
	}

}
