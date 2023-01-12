package DSA

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

func SignUp(userName string, money float64, nonce uint, userType string, userNum int, reputation int) Account {
	priKeyHash, _ := crypto.GenerateKey() //两处PublicKey意义不同
	fmt.Println(priKeyHash)
	//priKey, err := crypto.HexToECDSA(priKeyHash),此处是导入PrivateKey
	//if err != nil {
	//	panic(err)
	//}
	priKeyBytes := crypto.FromECDSA(priKeyHash)
	priKey := hex.EncodeToString(priKeyBytes)
	fmt.Printf("私钥为: %s\n", priKey)

	pubKey := priKeyHash.Public().(*ecdsa.PublicKey)
	// 获取公钥并去除头部0x04，原公钥65字节
	pubKeyBytes := crypto.FromECDSAPub(pubKey)[1:]
	fmt.Printf("公钥为: %s\n", hex.EncodeToString(pubKeyBytes))

	// 获取地址
	addr := crypto.PubkeyToAddress(*pubKey)
	fmt.Printf("地址为: %s\n", addr.Hex())
	return Account{
		publicKey:  hex.EncodeToString(pubKeyBytes),
		privateKey: priKey,
		Address:    addr.Hex(),
		UserName:   userName,
		UserType:   userType,
		UserNum:    userNum,
		Money:      money,
		Nonce:      nonce,
		TxPool:     nil,
	}
}
