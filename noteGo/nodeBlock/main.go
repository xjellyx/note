package main

import (
	"fmt"
	"github.com/noteGo/noteGo/nodeBlock/IntroductionBlock/block"
	"github.com/noteGo/noteGo/nodeBlock/IntroductionBlock/blockchain"
	"time"
)

func main()  {
	var(
		bc *blockchain.BlockChain
	)
	bc=new(blockchain.BlockChain)
	bc=blockchain.NewBlockChainInit()
	bc.AddNewBlock(&block.Block{TimeStamp:time.Now().Unix(),
		Data:[]byte("the second block"),})
	time.Sleep(3*time.Second)
	bc.AddNewBlock(&block.Block{TimeStamp:time.Now().Unix(),
		Data:[]byte("the third block"),})
	for _,v:=range bc.Blocks{
		fmt.Printf("timeStamp:%d\n",v.TimeStamp)
		fmt.Printf("Data:%s\n",v.Data)
		fmt.Printf("prev.hash:%x\n",v.PrevBlockHash)
		fmt.Printf("%x\n",v.Hash)
	}
}
