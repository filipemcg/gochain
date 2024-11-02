package main

import (
	"fmt"

	gochain "github.com/filipemcg/gochain/pkg"
)

func main() {
	// just so I don't forget that I can recover from panics
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	chain, err := gochain.NewChain()
	if err != nil {
		panic(err)
	}

	defer chain.Close()

	go RunApi(chain)

	select {}
}
