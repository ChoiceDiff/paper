package Consensus

type Account struct {
	publicKey  []byte
	privateKey []byte
	Address    []byte
	Name       string
	Money      float64
	Nonce      uint //交易次数
}
