package DSA

type UTXO map[string]map[int]int //商品剩余个数，用户名->商品编号->个数

//func NewUTXO() *UTXO {
//	return &UTXO{
//		UnspentTXOutput: make(map[string]map[int]int),
//	}
//}
