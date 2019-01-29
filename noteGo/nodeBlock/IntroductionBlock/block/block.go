package block

import (
	"bytes"
	"crypto/sha256"
	"strconv"
)

type Block struct {
	TimeStamp int64  // 时间截
	Data   []byte   // 区块链所记载的数据
	PrevBlockHash  []byte  // 上一区块哈希
	Hash []byte      // 本区块哈希
}

// 声明区块和处理上一个区块hash
func (b *Block)SetHash()(ret *Block){
	ret=new(Block)
	timeStamp:=[]byte(strconv.FormatInt(b.TimeStamp,10))
	headers:=bytes.Join([][]byte{b.PrevBlockHash,b.Data,timeStamp},[]byte{})
	hash:=sha256.Sum256(headers)
	b.Hash=hash[:]
	return
}

// 新建一个区块
func (b *Block)NewBlock()(ret *Block)  {
	b.SetHash()
	ret=new(Block)
	ret=b
	return
}
