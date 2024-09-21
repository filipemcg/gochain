package main

import (
	"fmt"

	gochain "github.com/filipemcg/gochain/pkg"
)

func main() {
	blockChain := []gochain.Block{}

	genisesPrevBytes := [32]byte{0}

	genisesBlock := gochain.NewBlock(1, "", genisesPrevBytes)

	blockChain = append(blockChain, *genisesBlock)

	for _, b := range blockChain {
		fmt.Printf("Block{Block#: %d, Data: %s, Nonce: %d, Prev: %x, Hash: %x}\n", b.Number, b.Data, b.Nonce, b.Prev, b.Hash)
	}
}
