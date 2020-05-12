package main

import (
	"github.com/gin-gonic/gin"
	"github.com/theiny/goldie-blocking-chain/adding"
	"github.com/theiny/goldie-blocking-chain/blockchain"
	"github.com/theiny/goldie-blocking-chain/listing"
	"github.com/theiny/goldie-blocking-chain/server"
	log "github.com/theiny/slog"
)

func main() {

	log.Info("Starting app...")

	s := server.New()
	bc := blockchain.InitBlockChain()

	s.Adding = adding.NewService(bc)
	s.Listing = listing.NewService(bc)

	gin.SetMode(gin.ReleaseMode)
	s.Router = gin.Default()

	s.Router.POST("/blockchain/add", s.AddBlock)
	s.Router.GET("blockchain/get", s.GetBlocks)

	s.Router.Run(":8080")
}
