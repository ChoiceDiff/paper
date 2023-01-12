package Consensus

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"math/rand"
	"paper/DSA"
	"strconv"
	"time"
)

var AllAccount DSA.BalanceTree

var AllUserAccount DSA.BalanceTree   //一般账户列表
var AllVerifyAccount DSA.BalanceTree //验证账户列表

var muptiUserAccountStateReplicas map[string]DSA.BalanceTree //addr->state
var muptiUserBlockchainReplicas map[string]DSA.BlockChain    //addr->blockchain

var Roles [7]string        //role types
var UserNum map[string]int //number of users
var Utxo DSA.UTXO          //unspent transaction output
var Products []DSA.Product //product list

var TxBroadCast map[string]chan DSA.Transaction //channel used to broadcast txs
var BlockBroadCast map[string]chan DSA.Block    //channel used to broadcast blocks

var Result *excelize.File

func init() {
	muptiUserAccountStateReplicas = make(map[string]DSA.BalanceTree)
	muptiUserBlockchainReplicas = make(map[string]DSA.BlockChain)

	UserNum = make(map[string]int) //访问map需要初始化
	TxBroadCast = make(map[string]chan DSA.Transaction)
	BlockBroadCast = make(map[string]chan DSA.Block)
	Roles = [7]string{"Supplier", "Business", "Retail", "Client", "Pay", "Express", "Factor"}

	Result := excelize.NewFile()
	// 创建一个工作表
	index := Result.NewSheet("Sheet1")
	// 设置工作簿的默认工作表
	Result.SetActiveSheet(index)
	// 根据指定路径保存文件
	if err := Result.SaveAs("Book1.xlsx"); err != nil {
		println(err.Error())
	}

	UserNum["Supplier"] = 5
	UserNum["Business"] = 15
	UserNum["Retail"] = 5
	UserNum["Client"] = 20

	//负责验证的节点成环
	UserNum["Pay"] = 2
	UserNum["Express"] = 2
	UserNum["Factor"] = 2

	rand.Seed(time.Now().UnixNano()) //设置随机数种子
	for i := 0; i < 10; i++ {
		price := 1 + 15*rand.Float64() //随机价格
		Products = append(Products, DSA.Product{
			Id:    uint(i),
			Price: price,
		})
	} //添加十件商品

	for i := 0; i < 7; i++ {
		rep := 0
		switch i {
		case 0:
			rep = 500
		case 1:
			rep = 500
		case 2:
			rep = 500
		case 3:
			rep = 100
		}
		for j := 0; j < UserNum[Roles[i]]; j++ { //注册用户
			userType := Roles[i]
			userNum := j
			name := userType + strconv.Itoa(userNum) //用户名
			newAccount := DSA.SignUp(name, 100000, 0, userType, j, rep)

			if i < 4 {
				AllUserAccount = append(AllUserAccount, newAccount) //账户余额默认100000
			} else {
				AllVerifyAccount = append(AllVerifyAccount, newAccount) //账户余额默认100000
			}
			AllAccount = append(AllAccount, newAccount)

			//每个地址（节点）维护一个副本
			addr := newAccount.Address

			muptiUserBlockchainReplicas[addr] = //区块链副本
				DSA.BlockChain{
					DSA.Block{
						Height:    0,
						TimeStamp: time.Now().String(),
						CoinBase:  [64]byte{0000000000000000000000000000000000000000000000000000000000000000},
						PrevHash:  [64]byte{0000000000000000000000000000000000000000000000000000000000000000},
						TxHash:    [64]byte{0000000000000000000000000000000000000000000000000000000000000000},
						Txs:       nil,
						ExtraData: "This is GENESIS BLOCK, watch out!!!",
					}, //need comma
				}
			TxBroadCast[name] = make(chan DSA.Transaction, 1024) //加入专属交易管道
			BlockBroadCast[name] = make(chan DSA.Block, 10)      //加入专属区块管道
			for k := 0; k < 10; k++ {
				goods := make(map[int]int)
				goods[k] = 100000
				Utxo[name] = goods
			} //每人有每种商品100000
		}
	}

	for _, account := range AllAccount { //状态树副本
		muptiUserAccountStateReplicas[account.Address] = AllAccount
	}

	//验证组成环，未进行优化！！！
	sizeVerify := len(AllVerifyAccount)
	for i, account := range AllVerifyAccount {
		account.Neighbour = &AllVerifyAccount[(i+1)%sizeVerify]
	}
}

func ConCmd() {
	for _, account := range AllUserAccount {
		go UserListenThread(account) //每个用户开启一个线程负责监听！！！
	}
	for _, account := range AllVerifyAccount {
		go VerifyGroupThread(account) //每个用户开启一个线程负责监听！！！
	}
	select {}
}

func UserListenThread(account DSA.Account) { //用户线程
	user := account.UserName
	fmt.Println("Start " + user + " Reputation: " + strconv.Itoa(account.Reputation))

	go ReceivedVerifyAddTxThread(account) //开启接收交易线程
	go CreateTxThread(account)
	go ListenAndCreateBlockAndMinerQualification(account)
	select {} //阻塞线程
}

func VerifyGroupThread(account DSA.Account) {
	user := account.UserName
	fmt.Println("Start " + user + "VerifyGroup!!!")

	go ReceivedVerifyAddTxThread(account)                 //开启接收交易线程，接收验证
	go ListenAndCreateBlockAndMinerQualification(account) //接收验证
}

func ListenAndCreateBlockAndMinerQualification(account DSA.Account) {

}

func CreateTxThread(account DSA.Account) {

}

func ReceivedVerifyAddTxThread(account DSA.Account) {
	name := account.UserName
	fmt.Println("!!!!!!!!!", name, "Start Receiving TX!!! ")
	for {
		newestTx := <-TxBroadCast[name]
		fmt.Println(name, " Received TX!!! ", newestTx)
	}
}

func UserThread(account DSA.Account) {
	for { //
		rand.Seed(time.Now().UnixNano())
		timeRand := rand.Intn(7) + 3 //3-10s一笔交易
		time.Sleep(time.Duration(timeRand) * time.Second)

		newTx := ChooseReceiverRandomlyAndCreateTx(routineIdOrTxSender, kind)
		fmt.Println(newTx)
		//加入交易池
		for userName, channel := range TxBroadCast {
			if userName != user {
				channel <- newTx
			}
		} //全部广播!!!
	}

	receiverId := uid
	receiverKind := kind
	for receiverKind == kind && receiverId == uid {
		rand.Seed(time.Now().UnixNano())
		randomKindId := rand.Intn(4)                          //后面3类不参与交易
		receiverKind = Roles[randomKindId]                    //随机挑一类
		receiverId = rand.Intn(10000) % UserNum[receiverKind] //每一类的人数不一样
		// panic: invalid argument to Intn

	} //随机选一个非自己的用户做交易
	rand.Seed(time.Now().UnixNano())
	sender := kind + strconv.Itoa(uid)
	receiver := receiverKind + strconv.Itoa(receiverId)
	productId := rand.Intn(1000) % 10 //随机选择一件商品
	num := 1 + rand.Intn(10000)%10    //随机选择数量
	price := Products[productId].Price * float64(num)
	return DSA.Transaction{
		Sender:    sender,
		Receiver:  receiver,
		ProductId: uint(productId),
		Num:       uint(num),
		Price:     price,
		Timestamp: time.Now().Unix(),
	}
}
