package main

import (
	"log"
	"powblockchain/wallet"
)

func init() {
	log.SetPrefix("Blockchain ")
}

func main() {
	myWallet := wallet.NewWallet()
	log.Print(myWallet.PrivateKey())
	log.Printf(myWallet.PrivateKeyString())
	log.Print(myWallet.PublicKey())
	log.Printf(myWallet.PublicKeyString())

	// myBlockChainAddress := "MY_BLOCKCHAIN_ADDRESS"
	// blockchain := NewBlockchain(myBlockChainAddress)

	// blockchain.AddTransaction("a", "b", 10.22)
	// blockchain.AddTransaction("a", "b", 2.4)
	// blockchain.Mine()
	// blockchain.Print()

	// blockchain.AddTransaction("c", "d", 42.1)
	// blockchain.AddTransaction("d", "c", 3.14)
	// blockchain.Mine()
	// blockchain.Print()

	// log.Printf("Total for A: %.1f", blockchain.CalculateTransactionTotal("a"))
	// log.Printf("Total for B: %.1f", blockchain.CalculateTransactionTotal("b"))
	// log.Printf("Total for miner: %.1f", blockchain.CalculateTransactionTotal(myBlockChainAddress))
}
