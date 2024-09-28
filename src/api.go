package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	gochain "github.com/filipemcg/gochain/pkg"
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

func JSON(w http.ResponseWriter, s interface{}) {
	b, err := json.Marshal(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func getBlock(w http.ResponseWriter, req *http.Request) {
	hash := req.PathValue("hash")
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
			JSON(w, response)
			return
		}
	}
	http.Error(w, "block not found", http.StatusNotFound)
}

func postBlock(w http.ResponseWriter, req *http.Request) {
	var requestBody struct {
		Data string `json:"data"`
	}

	if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
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
	JSON(w, response)
}

func RunApi() {
	mux := http.NewServeMux()

	genisesPrevBytes := [32]byte{0}

	genisesBlock := gochain.NewBlock(1, "", genisesPrevBytes)

	blockChain = append(blockChain, genisesBlock)

	mux.HandleFunc("GET /blocks/{hash}", getBlock)
	mux.HandleFunc("POST /blocks", postBlock)

	fmt.Println("Starting server on :9000")
	http.ListenAndServe(":9000", mux)
}
