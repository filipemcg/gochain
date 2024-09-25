package gochain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
)

type Block struct {
	Number int
	Nonce  int
	Data   string
	Prev   [32]byte
}

func (b *Block) toBytes() ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(b)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (b *Block) Hash() [32]byte {
	bytesArray, err := b.toBytes()
	if err != nil {
		panic(err)
	}
	return sha256.Sum256(bytesArray)
}

func NewBlock(blockNumber int, data string, prev [32]byte) Block {
	var nonce int
	var hash [32]byte
	var newBlock Block
	for {
		newBlock = Block{blockNumber, nonce, data, prev}
		bytesArray, err := newBlock.toBytes()
		if err != nil {
			panic(err)
		}
		hash = sha256.Sum256(bytesArray)
		if hash[0] == 0x00 && hash[1] == 0x00 {
			break
		}
		nonce++
	}
	return newBlock
}
