package defines

const (
	BUSINESS_UNKNOWN uint8 = iota
	BUSINESS_P2P_SWAP
)

const (
	EVENT_UNKNOWN uint32 = iota
	EVENT_P2P_SWAP_REQUIRE
	EVENT_P2P_SWAP_CONFIRM
)

const (
	CHAIN_CODE_UNKNOWN uint8 = iota
	CHAIN_CODE_BTC
	CHAIN_CODE_BCH
	CHAIN_CODE_ETH
)

const (
	CHAIN_BTC = "btc"
	CHAIN_BCH = "bch"
	CHAIN_ETH = "eth"

	BUSINESS_P2P_SWAP_NAME = "p2p"
)
