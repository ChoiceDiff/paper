package DSA

type TX struct {
	Sender    string
	Receiver  string
	Timestamp int64
	ProductId uint
	Num       uint
	Price     float64
}

type Transaction struct {
	Tx         TX
	TxHash     [32]byte
	PaySig     [32]byte
	ExpressSig [32]byte
	FactorSig  [32]byte
}

func NewTx(tx *TX) *Transaction {
	return &Transaction{
		Tx:         *tx,
		TxHash:     [32]byte{},
		PaySig:     [32]byte{},
		ExpressSig: [32]byte{},
		FactorSig:  [32]byte{},
	}
}

func VerifyTx() {

}
