package gochain

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Number int
	Nonce  int
	Data   string
	Hash   [32]byte
}

func (b *Block) String() string {
	return fmt.Sprintf("Block{Data: %s, Nonce: %d, Hash: %x}", b.Data, b.Nonce, b.Hash)
}

func (b *Block) CalculateHash() [32]byte {
	data := fmt.Sprintf("%d%s", b.Nonce, b.Data)
	return sha256.Sum256([]byte(data))
}
func NewBlock(blockNumber int, nonce int, data string) *Block {
	blockData := fmt.Sprintf("%d%d%s", blockNumber, nonce, data)
	hash := sha256.Sum256([]byte(blockData))

	return &Block{blockNumber, nonce, data, hash}
}
