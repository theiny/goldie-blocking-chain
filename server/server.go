package server

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/theiny/slog"
	"github.com/theiny/goldie-blocking-chain/adding"
	"github.com/theiny/goldie-blocking-chain/listing"
)

type data struct {
	Data string `json:"data"`
}

type server struct {
	Router     *gin.Engine
	Adding adding.Service
	Listing listing.Service
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

	s.Adding.AddBlock([]byte(req.Data))

	c.JSON(http.StatusOK, "Added new block")
}

func (s *server) GetBlocks(c *gin.Context) {
	bc := s.Listing.GetBlockChain()
	c.JSON(http.StatusOK, bc)
}