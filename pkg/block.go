package gochain

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Data string
	Hash [32]byte
}

func (b *Block) String() string {
	return fmt.Sprintf("Block{Data: %s, Hash: %x}", b.Data, b.Hash)
}

func NewBlock(data string) *Block {
	hash := sha256.Sum256([]byte(data))

	return &Block{data, hash}
}
