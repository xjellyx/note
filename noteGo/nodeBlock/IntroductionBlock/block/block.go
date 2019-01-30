package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	Index int64
	TimeStamp string  // 时间截
	Data   []byte   // 区块链所记载的数据
	BPM    int64
	PrevBlockHash  []byte  // 上一区块哈希
	Hash []byte      // 本区块哈希
}

// 声明区块和处理上一个区块hash
func (b *Block)SetHash()(ret *Block){
	ret=new(Block)
	b.TimeStamp=time.Now().String()
	timeStamp:=[]byte(b.TimeStamp)
	headers:=bytes.Join([][]byte{b.PrevBlockHash,b.Data,timeStamp,[]byte(strconv.FormatInt(b.Index,10)),
		[]byte(strconv.FormatInt(b.BPM,10))},[]byte{})
	hash:=sha256.Sum256(headers) // 转为hash
	b.Hash=hash[:]
	return
}
// 计算哈希值和set原理一样
func (n *Block)CalculateHash()(ret string)  {
	record:= string(n.Index) + string(n.TimeStamp) + string(n.BPM) + string(n.PrevBlockHash)
	h:=sha256.New()
	h.Write([]byte(record))
	hashed:=h.Sum(nil)
	n.Hash=[]byte(hex.EncodeToString(hashed))
	ret=hex.EncodeToString(hashed)
	return
}
// 新建一个区块
func (b *Block)NewBlock()(ret *Block)  {
	b.SetHash()
	ret=new(Block)
	ret=b
	return
}
