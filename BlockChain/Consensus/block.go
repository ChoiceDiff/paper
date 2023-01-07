package Consensus

type Block struct {
	Height    uint
	TimeStamp string
	CoinBase  []byte //miner
	PrevHash  []byte
	TxHash    []byte
	Txs       []Transaction
}
