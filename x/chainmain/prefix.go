package chainmain

const (
	CoinType       = 394
	FundraiserPath = "44'/394'/0'/0/1"
)

var (
	AccountAddressPrefix   = "cro"
	AccountPubKeyPrefix    = "cropub"
	ValidatorAddressPrefix = "crocncl"
	ValidatorPubKeyPrefix  = "crocnclpub"
	ConsNodeAddressPrefix  = "crocnclcons"
	ConsNodePubKeyPrefix   = "crocnclconspub"
	HumanCoinUnit          = "cro"
	BaseCoinUnit           = "basecro" // 10^-8 AKA "carson"
	CroExponential         = 8         // 10^8
	OneCroInCarson         = 1_0000_0000
)
