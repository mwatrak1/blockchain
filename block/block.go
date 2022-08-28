package block

import (
	"crypto/sha256"
	"encoding/json"
	"log"
	"powblockchain/transaction"
	"time"
)

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []*transaction.Transaction
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*transaction.Transaction) *Block {
	return &Block{
		timestamp:    time.Now().UnixNano(),
		nonce:        nonce,
		previousHash: previousHash,
		transactions: transactions,
	}
}

func (block *Block) Print() {
	log.Printf("timestamp			%d\n", block.timestamp)
	log.Printf("nonce				%d\n", block.nonce)
	log.Printf("previous hash			%x\n", block.previousHash)
	log.Printf("transactions:\n")
	for _, t := range block.transactions {
		t.Print()
	}
}

func (block *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Nonce        int                        `json:"nonce"`
		PreviousHash [32]byte                   `json:"previous_hash"`
		Timestamp    int64                      `json:"timestamp"`
		Transations  []*transaction.Transaction `json:"transactions"`
	}{
		Nonce:        block.nonce,
		PreviousHash: block.previousHash,
		Timestamp:    block.timestamp,
		Transations:  block.transactions,
	})
}

func (block *Block) Hash() [32]byte {
	m, _ := json.Marshal(block)
	return sha256.Sum256([]byte(m))
}
