Consensus Mechanism in Consortium Blockchain based on Multi-Role Reputation:

1.Seven Types of Role
    type BlockChainAccount struct {
    	gorm.Model
    	PublicKey           []byte  `json:"public_key"`
    	PrivateKey          []byte  `json:"private_key"`
    	Address             []byte  `json:"address"`
    	Name                string  `json:"name"`
    	Role                byte    `json:"role"`
    	TxNonce             uint64  `json:"nonce"` //tx counter
    	CorrectTxNonce      uint64
    	//ContinuousCorrectTxNumber uint, bonus
    	BlockNonce          uint64
    	CorrectBlockNonce   uint64
    	//ContinuousCorrectBlockNumber   uint,bonus
    	ContinuousAbsenceNonce        uint64//
    	CandidateRecent12Rounds [12]bool//sum>6, back to 500R
    	ContinuousAbsenceNonce uint64

    	Money               float64 `json:"money"`
    	Reputation          float64   `json:"reputation"` //user(node, if node==user), reputation value
    }//Supplier, Business, Customer(Retail, Client)

    type BlockChainAccount struct {
    	gorm.Model
    	PublicKey  []byte  `json:"public_key"`
    	PrivateKey []byte  `json:"private_key"`
    	Address    []byte  `json:"address"`
    	Name       string  `json:"name"`
    	Role       byte    `json:"role"`
    	Reputation float64   `json:"reputation"` //user(node, if node==user), reputation value
    }//Factor, Express, Pay

2.Transactions
    2.1 TX
    type Transaction struct {
    	gorm.Model
    	Sender   []byte
    	Receiver []byte
    	Product  Product
    	Num      uint
    	Money    float64
    	Type     byte
    }
    2.2 Transaction Type
        goodsBase(newGoods)
        level 1
        level 2

        if orderType!=Transfer{
            if receiver==Supplier{
                TxType=Super
            }else{
                TxType=Middle
            }
        }else {
            TxType1=SuperTransfer
            TxType2=MiddleTransfer
        }

3.Block Qualification
    BlockSizeLimit: txsNumbers 32 per block

    type Block struct {
    	Height    uint
    	TimeStamp string
    	PrevHash []byte
    	Txs       []Transaction
    }

    3.1 Initial Reputation Value
        Supplier:500 R

        Business:500 R
        Customer(Retail):500 R
        //按比例4:1

        Customer(Client):100 R

        Factor:100 R
        Express:100 R
        Pay:100 R

    3.2 Miner and NodeType(M:Miner, V:Verify)
        Supplier: M, V
        Business: M, V
        Customer(Retail): M, V
        Customer(Client): V

        Factor:V, S
        Express:V, S
        Pay:V, S

    3.3 Job && Block Packing
        Suppliers pack their Super&&SuperTransfer txs.
        Businesses and Retails pack their Middle&&MiddleTransfer txs.
        Clients sometimes are just senders who are responsible for verify the txs.

        Factor:Signature, Verify
        Express:Signature, Verify
        Pay:Signature, Verify

4.Consistency Process:
    4.1 UTXO
        unspentOutPuts: products, money
    4.2 Broadcast
        overlay, p2p, TCP
    4.3 Verify
        4.3.1 Txs
            UTXO enough(products, money)
        4.3.2 Block
            prevHash, txs
    4.4 Execute Txs in Block
        Update the states of nodes so that accounts will read up-to-date state from according node.

5. Reputation
        5.1 Block or Tx
            Punishment=0
            Proportion
            if BlockSucceed{
                SuperMiner+=(blockSize/blockLimit)*15R
                MiddleMiner+=(blockSize/blockLimit)*25R//Proportion can be changed.
            }
            if Fork{//Including wrong PrevHash is regarded as "Forking Attack".
                MinerWrongBlockTime+=1;//how to use?
                Punishment+=(blockSize/blockLimit)*BLOCKREWARD*pow(1.1,MinerForkingTime)//
                //1.1pow times punishment, because correct txs should have been uploaded in standard time.
                Miner-=Punishment
                }

            if HasWrongTxsInBlock{
                Punishment+=pow(1.1,wrongTXSTime)*BLOCKREWARD*(wrongTxsSize/blockLimit)
                Supplier-=pow(1.1,wrongTXSTime)*(1+wrongTxsSize/blockLimit)*5
                Client-=pow(1.1,wrongTXSTime)*(1+wrongTxsSize/blockLimit)*2
                Retail||Business-=pow(1.1,wrongTXSTime)*(1+wrongTxsSize/blockLimit)*3
                Signature-=pow(1.1,wrongTXSTime)*(1+wrongTxsSize/blockLimit)*1
            }//Multi-punishment if multi-errors. 剔除错误交易？

            if TxWasVerifiedWrong{
                Supplier-=pow(1.1,wrongTXSTime)*(1+wrongTxsSize/blockLimit)*2
                Client-=pow(1.1,wrongTXSTime)*(1+wrongTxsSize/blockLimit)*1
                Retail||Business-=pow(1.1,wrongTXSTime)*(1+wrongTxsSize/blockLimit)*2
                Signature-=pow(1.1,wrongTXSTime)*(1+wrongTxsSize/blockLimit)*1
            }//Multi-punishment if multi-errors. 剔除错误交易？

                Factor,Express,Pay-=1*pow(1.1,wrongSigTime)

            if txSucceed{
                Suppliers: 2 R
                Retails and Businesses+=2 R
                Client+=1 R
            }

            if verifyOrSignature{
                Repu+=0.2 R
            }
        5.2 Choosing Miner
            MinerStandard: in turn, Super/Middle 1:4 Alternately
            Sort Algorithm: 0.6*Reputation+0.2*100*correctBlockRate+0.2*100*correctTxRate*TxReward, first 5
            5 Candidates: 5 min timeout
        5.3 Fade Away
            Reputation-=(blockSize/blockLimit)*BLOCKREWARD*pow(1.1,ContinuousAbsenceNonce)

            per Day: Miners Reputation-=10
                     Pay, Express, Factor-=5//no more than 115
                     Client-=0.5//no more than 105

            if wrongBlockRate>=1/3||wrongTxRate>=1/3{
                Reputation=300
            }
        5.3 Oligarchy
            CandidateRecent12Rounds [12]bool//sum>6, back to 500R
        5.4 Force Quit
            if Initial Reputation==500 && ReputationNow<300{
                ForceQuit
            }
            if Client Reputation<60{//90 days(3 months) without contribution
                ForceQuit
            }
            if VerifyReputation<80{//No contributions during continuous more than 5 days. Rest at most 7 days.
                ForceQuit
            }

            You need to sign in by offline way.