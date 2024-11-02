package gochain

import (
	"encoding/hex"
	"fmt"
	"log"
)

var Chain *KV

func GenisesBlock() Block {
	return NewBlock(1, "Genises", [32]byte{0})
}

func NewChain() (*KV, error) {
	path := "./data"

	log.Printf("Using database at %s", path)
	kv, err := NewBadgerDb(path)
	if err != nil {
		panic(err)
	}
	defer kv.Close()

	Chain = kv

	genisesBlock := NewBlock(1, "Genises", [32]byte{0})
	genisesHash := genisesBlock.Hash()
	tt := hex.EncodeToString(genisesHash[:])

	fmt.Println(genisesHash[:])
	fmt.Println(tt)

	genisesHashBytes, _ := hex.DecodeString(tt)

	genises, err := Chain.Get(genisesHashBytes)
	if err != nil {
		genesisByteBlock, error := genisesBlock.toBytes()
		if error != nil {
			return nil, error
		}
		Chain.Set(genisesHash[:], genesisByteBlock)
		genises, _ := Chain.Get(genisesHash[:])
		fmt.Println(BlockFromBytes(genises))
		return kv, nil
	}

	fmt.Println(BlockFromBytes(genises))
	return kv, nil
}

func PrintChain() {
	fmt.Println("Printing chain")

}
