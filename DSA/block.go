package DSA

type Block struct {
	Height    uint
	TimeStamp string
	CoinBase  [64]byte //miner
	PrevHash  [64]byte
	TxHash    [64]byte //Merkle Root
	Txs       []Transaction
	ExtraData string
}

func NewBlock(msgSender [64]byte, packed []Transaction, previousHash [64]byte) *Block { //接收参数
	return &Block{
		Height:    0,
		TimeStamp: "",
		CoinBase:  [64]byte{},
		PrevHash:  [64]byte{},
		TxHash:    [64]byte{},
		Txs:       nil,
		ExtraData: "",
	}
}
