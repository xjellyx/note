package blockchain

import (
	"github.com/noteGo/noteGo/nodeBlock/IntroductionBlock/block"
	"time"
)

// 实现一个区块链

// 声明一个区块链结构体
type BlockChain struct {
	Blocks []*block.Block  // 保存区块的信息切片
}

// 先创世一个区块
func GenesisBlock() *block.Block {
	var(
		blo *block.Block
	)
	blo=new(block.Block)
	blo.TimeStamp=time.Now().Unix()
	blo.Data=[]byte("my first block")
	blo.PrevBlockHash=[]byte{}
	return blo.NewBlock()
}

// 把区块添加到区块链中
func (bc *BlockChain)AddNewBlock(b *block.Block)  {
	b.PrevBlockHash=bc.Blocks[len(bc.Blocks)-1].Hash // 获取上一块区块的hash
	newBlock:=b.NewBlock()   // 声明新的区块
	bc.Blocks=append(bc.Blocks,newBlock)
}

// 初始化区块链,也就是生成第一条区块链
func NewBlockChainInit() *BlockChain  {
	return &BlockChain{[]*block.Block{GenesisBlock()}}
}

// 工厂模式
func (bc *BlockChain)FactoryBlocks()(ret []*block.Block)  {
	ret=bc.Blocks
	return
}