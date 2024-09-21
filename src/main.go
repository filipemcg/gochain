package main

import (
	"fmt"

	gochain "github.com/filipemcg/gochain/pkg"
)

func main() {
	genisesBlock := gochain.NewBlock(1, 72608, "")
	fmt.Println(genisesBlock)
}
