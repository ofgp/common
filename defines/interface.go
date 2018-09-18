package defines

type PushEvent interface {
	GetBusiness() string
	GetEventType() uint32
	GetTxID() string
	GetHeight() int64
	GetTranxIx() int
	GetAmount() uint64
	GetFee() uint64
	GetFrom() uint32
	GetTo() uint32
	GetData() []byte
}
