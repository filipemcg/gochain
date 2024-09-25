package main

import (
	"fmt"
	"net/http"

	gochain "github.com/filipemcg/gochain/pkg"
	"github.com/gin-gonic/gin"
)

type Block struct {
	Number int    `json:"number"`
	Data   string `json:"data"`
	Nonce  int    `json:"nonce"`
	Prev   string `json:"prev"`
}

type Response struct {
	Block Block  `json:"block"`
	Hash  string `json:"hash"`
}

var blockChain = []gochain.Block{}

func getBlock(c *gin.Context) {
	hash := c.Param("hash")
	for _, block := range blockChain {
		if fmt.Sprintf("%x", block.Hash()) == hash {
			response := Response{
				Block: Block{
					Number: block.Number,
					Data:   block.Data,
					Nonce:  block.Nonce,
					Prev:   fmt.Sprintf("%x", block.Prev),
				},
				Hash: fmt.Sprintf("%x", block.Hash()),
			}
			c.JSON(http.StatusOK, response)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "block not found"})
}

func postBlock(c *gin.Context) {
	var requestBody struct {
		Data string `json:"data"`
	}
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	prev := blockChain[len(blockChain)-1].Hash()
	newBlock := gochain.NewBlock(len(blockChain)+1, requestBody.Data, prev)
	blockChain = append(blockChain, newBlock)

	response := Response{
		Block: Block{
			Number: newBlock.Number,
			Data:   newBlock.Data,
			Nonce:  newBlock.Nonce,
			Prev:   fmt.Sprintf("%x", newBlock.Prev),
		},
		Hash: fmt.Sprintf("%x", newBlock.Hash()),
	}
	c.JSON(http.StatusCreated, response)
}

func RunApi() {
	r := gin.Default()

	genisesPrevBytes := [32]byte{0}

	genisesBlock := gochain.NewBlock(1, "", genisesPrevBytes)

	blockChain = append(blockChain, genisesBlock)

	r.GET("/blocks/:hash", getBlock)
	r.POST("/blocks", postBlock)

	fmt.Println("Starting server on :9000")
	r.Run(":9000")
}
