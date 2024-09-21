package main

import (
	"fmt"

	gochain "github.com/filipemcg/gochain/pkg"
)

func main() {
	genisesPrevBytes := [32]byte{}
	copy(genisesPrevBytes[:], "00000000000000000000000000000000")
	genisesBlock := gochain.NewBlock(1, "", genisesPrevBytes)
	fmt.Println(genisesBlock)
}
