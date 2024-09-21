package main

import (
	"fmt"
	"net/http"

	gochain "github.com/filipemcg/gochain/pkg"
	"github.com/gin-gonic/gin"
)

var blockChain = []gochain.Block{}

func getBlock(c *gin.Context) {
	hash := c.Param("hash")
	for _, block := range blockChain {
		if fmt.Sprintf("%x", block.Hash) == hash {
			serializedBlock := map[string]interface{}{
				"Number": block.Number,
				"Data":   block.Data,
				"Nonce":  block.Nonce,
				"Prev":   fmt.Sprintf("%x", block.Prev),
				"Hash":   fmt.Sprintf("%x", block.Hash),
			}
			c.JSON(http.StatusOK, serializedBlock)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "block not found"})
}

func postBlock(c *gin.Context) {
	data := c.PostForm("data")
	prev := blockChain[len(blockChain)-1].Hash
	newBlock := gochain.NewBlock(len(blockChain)+1, data, prev)
	blockChain = append(blockChain, *newBlock)

	serializedBlock := map[string]interface{}{
		"Number": newBlock.Number,
		"Data":   newBlock.Data,
		"Nonce":  newBlock.Nonce,
		"Prev":   fmt.Sprintf("%x", newBlock.Prev),
		"Hash":   fmt.Sprintf("%x", newBlock.Hash),
	}
	c.JSON(http.StatusCreated, serializedBlock)
}

func main() {
	r := gin.Default()

	blockChain = []gochain.Block{}

	genisesPrevBytes := [32]byte{0}

	genisesBlock := gochain.NewBlock(1, "", genisesPrevBytes)

	blockChain = append(blockChain, *genisesBlock)

	r.GET("/blocks/:hash", getBlock)
	r.POST("/blocks", postBlock)

	fmt.Println("Starting server on :8080")
	r.Run(":8080")
}
