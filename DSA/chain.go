package DSA

func ExpandBlockChain() {
	listenBlock := make(chan Block)
	for {
		if newReceivedBlock, ok := <-listenBlock; ok == true {

		}
	}
}

//go Consensus.ExpandBlockChain()
