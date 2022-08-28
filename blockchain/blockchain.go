package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

const (
	MINING_DIFFICULTY    = 3
	MINING_REWARD_SENDER = "BLOCKCHAIN"
	MINING_REWARD        = 1.0
)

type Transaction struct {
	senderAddress    string
	recipientAddress string
	value            float32
}

func NewTransaction(recipientAddress string, senderAddress string, value float32) *Transaction {
	return &Transaction{
		senderAddress:    senderAddress,
		recipientAddress: recipientAddress,
		value:            value,
	}
}

func (transaction *Transaction) Print() {
	log.Printf("%s\n", strings.Repeat("-", 64))
	log.Printf("sender address			%s\n", transaction.senderAddress)
	log.Printf("recipient address		%s\n", transaction.recipientAddress)
	log.Printf("value				%.2f\n", transaction.value)
}

func (transaction *Transaction) MashalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderAddress    string  `json:"sender_address"`
		RecipientAddress string  `json:"recipient_address"`
		Value            float32 `json:"value"`
	}{
		SenderAddress:    transaction.senderAddress,
		RecipientAddress: transaction.recipientAddress,
		Value:            transaction.value,
	})
}

type Pool struct {
	transactions []*Transaction
}

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []*Transaction
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
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
		Nonce        int            `json:"nonce"`
		PreviousHash [32]byte       `json:"previous_hash"`
		Timestamp    int64          `json:"timestamp"`
		Transations  []*Transaction `json:"transactions"`
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

type Blockchain struct {
	transactionPool Pool
	chain           []*Block
	minerAddress    string
}

func NewBlockchain(minerAddress string) *Blockchain {
	firstBlock := &Block{}
	blockchain := &Blockchain{
		transactionPool: Pool{transactions: []*Transaction{}},
		chain:           []*Block{firstBlock},
		minerAddress:    minerAddress,
	}
	return blockchain
}

func (blockChain *Blockchain) CreateBlock(nonce int) *Block {
	lastBlock := blockChain.LastBlock()
	block := NewBlock(nonce, lastBlock.Hash(), blockChain.transactionPool.transactions)
	blockChain.chain = append(blockChain.chain, block)
	blockChain.transactionPool.transactions = []*Transaction{}
	return block
}

func (blockchain *Blockchain) LastBlock() *Block {
	return blockchain.chain[len(blockchain.chain)-1]
}

func (blockchain *Blockchain) AddTransaction(recipientAddress string, senderAddress string, value float32) *Transaction {
	transaction := NewTransaction(recipientAddress, senderAddress, value)
	blockchain.transactionPool.transactions = append(blockchain.transactionPool.transactions, transaction)
	return transaction
}

func (blockchain *Blockchain) ValidProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{timestamp: 0, nonce: nonce, previousHash: previousHash, transactions: transactions}
	guessBlockStringifiedHash := fmt.Sprintf("%x", guessBlock.Hash())
	return guessBlockStringifiedHash[:difficulty] == zeros
}

func (blockchain *Blockchain) ProofOfWork() int {
	transactions := blockchain.CopyTransactionPool()
	lastBlock := blockchain.LastBlock()
	previousHash := lastBlock.previousHash
	nonce := 0

	for !blockchain.ValidProof(nonce, previousHash, transactions, MINING_DIFFICULTY) {
		nonce += 1
	}
	return nonce
}

func (blockchain *Blockchain) Mine() bool {
	log.Println("Mining...")
	blockchain.AddTransaction(blockchain.minerAddress, MINING_REWARD_SENDER, MINING_REWARD)
	nonce := blockchain.ProofOfWork()
	blockchain.CreateBlock(nonce)
	return true
}

func (blockchain *Blockchain) CalculateTransactionTotal(blockchainAddress string) float32 {
	var totalAmount float32 = 0.0

	for _, block := range blockchain.chain {
		for _, transaction := range block.transactions {
			if transaction.recipientAddress == blockchainAddress {
				totalAmount += transaction.value
			}
			if transaction.senderAddress == blockchainAddress {
				totalAmount -= transaction.value
			}
		}
	}

	return totalAmount
}

func (blockchain *Blockchain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, transaction := range blockchain.transactionPool.transactions {
		transactions = append(transactions, NewTransaction(
			transaction.recipientAddress,
			transaction.senderAddress,
			transaction.value,
		))
	}
	return transactions
}

func (blockchain *Blockchain) Print() {
	for i, block := range blockchain.chain {
		log.Printf("%s Block %d %s\n", strings.Repeat("-", 32), i, strings.Repeat("-", 32))
		block.Print()
	}
}
