package serve

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/olefen/note/block/block"
	"github.com/olefen/note/block/blockchain"
	"net/http"
	"strconv"
)

var (
	BC      *blockchain.BlockChain
	DataMap = make(map[string]interface{})
)

// NewBlock 生成区块
func NewBlock(ctx *gin.Context) {
	d := ctx.PostForm("data")
	bpm := ctx.PostForm("bpm")

	var (
		data = new(block.Block)
	)
	b, err := strconv.ParseInt(bpm, 0, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	data.Data = d
	data.BPM = b
	if data, err = BC.NewBlock(data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	//
	DataMap[data.Hash] = data
	ctx.AbortWithStatusJSON(http.StatusOK, data)
}

// GetBlock 通过hash获取区块信息
func GetBlock(ctx *gin.Context) {
	hash := ctx.Param("hash")
	var (
		err error
	)
	if val, ok := DataMap[hash]; !ok {
		err = errors.New("不存在该区块")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	} else {
		ctx.AbortWithStatusJSON(200, val)
	}
}

// GetAllBlock 获取全部区块
func GetAllBlock(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusOK, BC)
}

// 生成创世区块
func init() {
	BC = new(blockchain.BlockChain)
	BC.Blocks = append(BC.Blocks, block.GenerateBlock())
	DataMap[BC.Blocks[0].Hash] = BC.Blocks[0]
}
