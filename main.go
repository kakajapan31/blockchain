package main

import "fmt"

func main() {
	bc := NewBlockChain()

	bc.AddBlock("Send 1 bitcoin to kaka")
	bc.AddBlock("Send 2 bitcoin to kazu")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
