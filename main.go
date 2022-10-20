package main

import (
	"flag"
	"powblockchain/node"
)

func main() {
	port := flag.Uint("port", 5000, "TCP Port for Blockchain Node")
	flag.Parse()

	blockchainNode := node.NewBlockchainNode(uint16(*port))
	blockchainNode.Run()
}
