package server

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	blockchain "github.com/theiny/goldie-blocking-chain/blockchain"
	log "github.com/theiny/slog"
)

type data struct {
	Data string `json:"data"`
}

type server struct {
	Router     *gin.Engine
	BlockChain *blockchain.BlockChain
}

func New() *server {
	return &server{}
}

func (s *server) AddBlock(c *gin.Context) {
	var req data
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		log.Error(err)
		return
	}

	s.BlockChain.AddBlock([]byte(req.Data))
}

func (s *server) GetBlocks(c *gin.Context) {
	c.JSON(http.StatusOK, s.BlockChain)
}