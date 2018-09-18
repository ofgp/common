package defines

type PushEvent interface {
	GetBusiness() string
	GetEventType() uint32
	GetTxID() string
	GetHeight() int64
	GetTranxIx() int
	GetAmount() uint64
	GetFee() uint64
	GetFrom() byte
	GetTo() byte
	GetData() []byte
}
