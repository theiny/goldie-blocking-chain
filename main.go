package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/theiny/goldie-blocking-chain/blockchain"
	"github.com/theiny/goldie-blocking-chain/listing"
	"github.com/theiny/goldie-blocking-chain/sending"
	"github.com/theiny/goldie-blocking-chain/server"
)

func main() {

	log.Println("Starting app...")

	s := server.New()

	// Init the blockchain using a default hardcoded address.
	bc := blockchain.InitBlockChain(blockchain.GenesisAddress)

	// create new services for sending and listing
	s.Sending = sending.NewService(bc)
	s.Listing = listing.NewService(bc)

	// load router
	gin.SetMode(gin.ReleaseMode)
	s.Router = gin.Default()

	// load routes
	s.LoadRoutes()

	s.Router.Run(":8080")
}
