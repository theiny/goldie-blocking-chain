package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transaction struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount int    `json:"amount"`
}

func (s *server) LoadRoutes() {
	bc := s.Router.Group("/api/v1/blockchain")
	{
		bc.GET("/list", s.ListBlockchain)
		bc.GET("/balance/:addr", s.ShowBalance)
		bc.POST("/send", s.SendGold)
	}
}

func (s *server) SendGold(c *gin.Context) {
	var tx transaction
	err := json.NewDecoder(c.Request.Body).Decode(&tx)
	if err != nil {
		respondErr(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = s.Sending.Send(tx.From, tx.To, tx.Amount)
	if err != nil {
		respondErr(c, http.StatusInternalServerError, err.Error())
		return
	}

	respond(c, fmt.Sprintf("%s successfully sent %d golden nuggets to %s", tx.From, tx.Amount, tx.To))
}

func (s *server) ListBlockchain(c *gin.Context) {
	c.JSON(http.StatusOK, s.Listing.GetBlockChain())
}

func (s *server) ShowBalance(c *gin.Context) {
	addr := c.Param("addr")
	if addr == "" {
		respondErr(c, http.StatusBadRequest, "Address is empty")
		return
	}

	balance := s.Listing.GetBalance(addr)

	respond(c, fmt.Sprintf("Balance of %s: %d", addr, balance))
}
