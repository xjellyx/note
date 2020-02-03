package blockchain

import (
	"errors"
	"github.com/olefen/note/block/block"
)

// 实现一个区块链

// 声明一个区块链结构体
type BlockChain struct {
	Blocks []*block.Block // 保存区块的信息切片
}

// BlockNew 新建一个区块
func (bc *BlockChain) NewBlock(b *block.Block) (ret *block.Block, err error) {
	if bc == nil {
		err = errors.New("block chain is nil")
		return
	}
	var (
		data *block.Block
	)
	b.PrevBlockHash = bc.Blocks[len(bc.Blocks)-1].Hash // 获取上一块区块的hash
	b.Index = bc.Blocks[len(bc.Blocks)-1].Index + 1
	data = block.NewBlock(b)

	// 添加区块到链中
	bc.AddBlock(data)

	//
	ret = data
	return
}

// AddBlock 添加区块到链中
func (bc *BlockChain) AddBlock(b *block.Block) {
	bc.Blocks = append(bc.Blocks, b)
}

// NewBlockChainInit 初始化区块链,也就是生成第一条区块链
func InitBlockChain(blo *block.Block) *BlockChain {
	return &BlockChain{[]*block.Block{block.NewBlock(blo)}}
}

// FactoryBlocks 工厂模式
func (bc *BlockChain) FactoryBlocks() (ret []*block.Block) {
	ret = bc.Blocks
	return
}

// ValidBlock 区块正确性验证, false 验证失败
func ValidBlock(oldBlock *block.Block, newBlock *block.Block) bool {
	// 上一个区块index不等于新的区块index-1,返回
	if oldBlock.Index != newBlock.Index-1 {
		return false
	}
	// 新区块的prev hash 与旧区块hash比较是否相同
	if string(oldBlock.Hash) != string(newBlock.PrevBlockHash) {
		return false
	}
	return true
}

// ReplaceChain 如果有两个节点，选择最长的那个节点
func (bc *BlockChain) ReplaceChain(newBlocks []*block.Block) {
	if len(newBlocks) > len(bc.Blocks) {
		bc.Blocks = newBlocks
	}

}
