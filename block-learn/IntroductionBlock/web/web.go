package web

import (
	"github.com/LnFen/note/block-learn/IntroductionBlock/block"
	"github.com/LnFen/note/block-learn/IntroductionBlock/blockchain"
	"github.com/gorilla/mux"

	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

var NewBlock *blockchain.BlockChain

func Run() (err error) {
	if err = http.ListenAndServe("127.0.0.1:8080", setMuxRouter()); err != nil {
		return
	}
	return
}

// 新写路由
func setMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/get/all", handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/post/l", handleWriteBlock).Methods("POST")
	return muxRouter
}

// 处理get请求
func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	var (
		bytes []byte
		err   error
	)
	if bytes, err = json.MarshalIndent(NewBlock.Blocks, "", ""); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

// 处理post请求
type Msg struct {
	Index int
	Data  string
	BPM   int
}

func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		msg *Msg
		blo *block.Block
	)
	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&msg); err != nil {
		log.Println(err.Error())
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()
	blo = new(block.Block)
	blo.BPM = int64(msg.BPM)
	blo.Index = int64(msg.Index)
	blo.Data = []byte(msg.Data)
	bloc := NewBlock.AddNewBlock(blo) // 添加新的区块
	if !blockchain.IsBlockValid(NewBlock.Blocks[len(NewBlock.Blocks)-1], bloc) {
		err = errors.New("the block is not same as previous")
		log.Fatal(err.Error())
		return
	} else {
		NewBlock.AddBlock(bloc)
	}

	respondWithJSON(w, r, http.StatusCreated, NewBlock)

}

// 处理
func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500:Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}
func init() {
	NewBlock = blockchain.NewBlockChainInit(&block.Block{TimeStamp: time.Now().String(),
		Data: []byte("the first block")})
}
