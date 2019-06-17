package blockchain

import (
	"github.com/srlemon/note/block_/IntroductionBlock/block"
)

// 实现一个区块链

// 声明一个区块链结构体
type BlockChain struct {
	Blocks []*block.Block // 保存区块的信息切片
}

// 先创世一个区块
func GenesisBlock(blo *block.Block) (ret *block.Block) {
	if blo == nil {
		return blo.NewBlock()
	} else {
		blo.Data = []byte("the first block")
		blo.Index = 0
		blo.BPM = 0
		blo.CalculateHash()
		ret = new(block.Block)
		ret = blo
		return
	}
	return
}

// 新建一个区块
func (bc *BlockChain) AddNewBlock(b *block.Block) (ret *block.Block) {
	b.PrevBlockHash = bc.Blocks[len(bc.Blocks)-1].Hash // 获取上一块区块的hash
	b.Index = bc.Blocks[len(bc.Blocks)-1].Index + 1
	newBlock := b.NewBlock() // 声明新的区块
	ret = new(block.Block)
	ret = newBlock
	return
}

// 添加区块到链中
func (bc *BlockChain) AddBlock(b *block.Block) {
	bc.Blocks = append(bc.Blocks, b)
}

// 初始化区块链,也就是生成第一条区块链
func NewBlockChainInit(blo *block.Block) *BlockChain {
	return &BlockChain{[]*block.Block{GenesisBlock(blo)}}
}

// 工厂模式
func (bc *BlockChain) FactoryBlocks() (ret []*block.Block) {
	ret = bc.Blocks
	return
}

// 区块正确性验证
func IsBlockValid(oldBlock *block.Block, newBlock *block.Block) bool {
	if oldBlock.Index != newBlock.Index-1 {
		return false
	}
	if string(oldBlock.Hash) != string(newBlock.PrevBlockHash) {
		return false
	}
	return true
}

// 如果有两个节点，当两个节点上的区块链长度不同时，我们应该选择哪条链呢？最简单的方法，选择较长的那个。
func (bc *BlockChain) ReplaceChain(newBlocks []*block.Block) {
	if len(newBlocks) > len(bc.Blocks) {
		bc.Blocks = newBlocks
	}

}
