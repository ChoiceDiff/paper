package Consensus

type UTXO struct {
	UnspentTXOutput map[string]map[int]int //商品剩余个数，用户名->商品个数
}
