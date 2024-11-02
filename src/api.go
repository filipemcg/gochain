package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/filipemcg/gochain/docs"
	gochain "github.com/filipemcg/gochain/pkg"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Block represents a block in the blockchain
type Block struct {
	Number int    `json:"number"`
	Data   string `json:"data"`
	Nonce  int    `json:"nonce"`
	Prev   string `json:"prev"`
}

// Response represents the response structure for a block
type Response struct {
	Block Block  `json:"block"`
	Hash  string `json:"hash"`
}

// ErrorResponse represents the error response structure
type ErrorResponse struct {
	Message string `json:"message"`
}

// PostBlockRequest represents the request body for creating a new block
type PostBlockRequest struct {
	Data string `json:"data"`
}

var blockChain = []gochain.Block{}
var kv *gochain.KV

// getBlock retrieves a block by its hash
// @Summary Get a block by hash
// @Description Get a block by its hash
// @Tags blocks
// @Param hash path string true "Block Hash"
// @Success 200 {object} Response
// @Failure 404 {object} ErrorResponse
// @Router /blocks/{hash} [get]
func getBlock(c *gin.Context) {
	hash := c.Param("hash")
	genisesHashBytes, _ := hex.DecodeString(hash)

	genises := gochain.GenisesBlock()
	genisesHash := genises.Hash()

	if bytes.Equal(genisesHash[:], genisesHashBytes) {
		fmt.Println("Genises block")
	}

	tt := hex.EncodeToString(genisesHash[:])
	ttt, _ := hex.DecodeString(tt)

	value, err := kv.Get(ttt)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "block not found"})
		return
	}

	block, err := gochain.BlockFromBytes(value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "error decoding block"})
		return
	}

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
}

// postBlock creates a new block
// @Summary Create a new block
// @Description Create a new block with the provided data
// @Tags blocks
// @Accept json
// @Produce json
// @Param block body PostBlockRequest true "Block Data"
// @Success 201 {object} Response
// @Failure 400 {object} ErrorResponse
// @Router /blocks [post]
func postBlock(c *gin.Context) {
	var requestBody PostBlockRequest
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request body"})
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

// RunApi starts the API server
func RunApi(chain *gochain.KV) {
	kv = chain

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	apiV1 := r.Group("/api/v1")

	apiV1.GET("/blocks/:hash", getBlock)
	apiV1.POST("/blocks", postBlock)

	fmt.Println("Starting server on :9000")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":9000")
}
