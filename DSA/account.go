package DSA

type Account struct {
	publicKey  string
	privateKey string
	Address    string
	UserName   string
	UserType   string
	UserNum    int
	Money      float64
	Nonce      uint //交易次数
	TxPool     []Transaction
	Reputation int
	Neighbour  *Account
}
