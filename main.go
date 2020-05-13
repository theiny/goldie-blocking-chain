package main

import (
	"github.com/gin-gonic/gin"
	"github.com/theiny/goldie-blocking-chain/blockchain"
	"github.com/theiny/goldie-blocking-chain/listing"
	"github.com/theiny/goldie-blocking-chain/sending"
	"github.com/theiny/goldie-blocking-chain/server"
	log "github.com/theiny/slog"
)

func main() {

	log.Info("Starting app...")

	s := server.New()
	bc := blockchain.InitBlockChain("Theiny")

	s.Sending = sending.NewService(bc)
	s.Listing = listing.NewService(bc)

	gin.SetMode(gin.ReleaseMode)
	s.Router = gin.Default()

	s.LoadRoutes()

	s.Router.Run(":8080")
}
