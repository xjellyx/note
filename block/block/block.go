package block

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index         int
	TimeStamp     int64  // 时间截
	Data          string // 区块链所记载的数据
	BPM           int64
	PrevBlockHash string // 上一区块哈希
	Hash          string // 本区块哈希
}

// NewBlock 新建一个区块
func NewBlock(in *Block) (out *Block) {
	if in != nil {
		out = in
	} else {
		out = new(Block)
	}
	var (
		now = time.Now()
	)

	if out.TimeStamp <= 0 {
		out.TimeStamp = now.Unix()
	}

	record := string(out.Index) + string(out.TimeStamp) + string(out.BPM) + out.PrevBlockHash
	h := sha256.New()
	h.Write([]byte(record))
	out.Hash = hex.EncodeToString(h.Sum(nil))

	return
}

// GenerateBlock 生成创世区块
func GenerateBlock() (ret *Block) {
	ret = new(Block)
	ret.TimeStamp = time.Now().Unix()
	ret.BPM = 0
	ret.Index = 1
	ret.Data = "创世区块!!!"
	record := string(ret.Index) + string(ret.TimeStamp) + string(ret.BPM) + ret.PrevBlockHash
	h := sha256.New()
	h.Write([]byte(record))
	ret.Hash = hex.EncodeToString(h.Sum(nil))
	return
}
