package main

import (
	"fmt"

	gochain "github.com/filipemcg/gochain/pkg"
)

func main() {
	genisesBlock := gochain.NewBlock("")
	fmt.Println(genisesBlock)
}
