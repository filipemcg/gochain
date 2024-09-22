package gochain

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Number int
	Nonce  int
	Data   string
	Prev   [32]byte
	Hash   [32]byte
}

func calculateNonce(blockNumber int, data string, prev [32]byte) int {
	var nonce int
	var hash [32]byte
	for {
		blockData := fmt.Sprintf("%d%d%s%x", blockNumber, nonce, data, prev)
		hash = sha256.Sum256([]byte(blockData))
		if hash[0] == 0x00 && hash[1] == 0x00 {
			break
		}
		nonce++
	}
	return nonce
}

func NewBlock(blockNumber int, data string, prev [32]byte) *Block {
	nonce := calculateNonce(blockNumber, data, prev)

	blockData := fmt.Sprintf("%d%d%s%x", blockNumber, nonce, data, prev)
	hash := sha256.Sum256([]byte(blockData))

	return &Block{blockNumber, nonce, data, prev, hash}
}
