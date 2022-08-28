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
	log.Println(myWallet.Address())
}
