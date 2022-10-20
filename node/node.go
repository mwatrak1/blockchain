package node

import (
	"io"
	"log"
	"net/http"
	"powblockchain/blockchain"
	"powblockchain/wallet"
	"strconv"
)

var cache map[string]*blockchain.Blockchain = make(map[string]*blockchain.Blockchain)

type BlockchainNode struct {
	port uint16
}

func NewBlockchainNode(port uint16) *BlockchainNode {
	return &BlockchainNode{
		port: port,
	}
}

func (blockchainNode *BlockchainNode) Port() uint16 {
	return blockchainNode.port
}

func (blockchainNode *BlockchainNode) GenerateUrl(basePath string) string {
	return basePath + ":" + strconv.Itoa(int(blockchainNode.Port()))
}

func (blockchainNode *BlockchainNode) GetBlockchain() *blockchain.Blockchain {
	bc, ok := cache["blockchain"]
	if !ok {
		minersWallet := wallet.NewWallet()
		bc = blockchain.NewBlockchain(minersWallet.Address(), blockchainNode.Port())
		cache["blockchain"] = bc
		log.Printf("Blockchain wallet registered with address %v", minersWallet.Address())
	}
	return bc
}

func (blockchainNode *BlockchainNode) GetChain(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		bc := blockchainNode.GetBlockchain()
		m, _ := bc.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		log.Printf("Error: Invalid HTTP Method")
	}
}

func (blockchainNode *BlockchainNode) Run() {
	http.HandleFunc("/", blockchainNode.GetChain)
	http.ListenAndServe(blockchainNode.GenerateUrl("0.0.0.0"), nil)
}
