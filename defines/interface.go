package defines

type PushEvent interface {
	GetBusiness() uint8
	GetEventType() uint32
	GetProposal() string
	GetTxID() string
	GetHeight() int64
	GetTranxIx() int
	GetAmount() uint64
	GetFee() uint64
	GetFrom() byte
	GetTo() byte
	GetData() []byte
}
