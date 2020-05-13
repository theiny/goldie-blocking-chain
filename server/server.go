package server

import (
	"github.com/gin-gonic/gin"
	"github.com/theiny/goldie-blocking-chain/listing"
	"github.com/theiny/goldie-blocking-chain/sending"
)

type server struct {
	Router  *gin.Engine
	Sending sending.Service
	Listing listing.Service
}

func New() *server {
	return &server{}
}
