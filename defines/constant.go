package defines

const (
	BUSINESS_UNKNOWN uint8 = iota
	BUSINESS_P2P_SWAP
	BUSINESS_COIN_MINT
)

const (
	EVENT_UNKNOWN uint32 = iota
	EVENT_P2P_SWAP_REQUIRE
	EVENT_P2P_SWAP_CONFIRM
	EVENT_COIN_MINT_REQUIRE
	EVENT_COIN_MINT_CONFIRM
)

const (
	CHAIN_CODE_UNKNOWN uint8 = iota
	CHAIN_CODE_BTC
	CHAIN_CODE_BCH
	CHAIN_CODE_ETH
)

const (
	CHAIN_NAME_BTC = "btc"
	CHAIN_NAME_BCH = "bch"
	CHAIN_NAME_ETH = "eth"

	BUSINESS_P2P_SWAP_NAME  = "p2p"
	BUSINESS_COIN_MINT_NAME = "mint"
)
