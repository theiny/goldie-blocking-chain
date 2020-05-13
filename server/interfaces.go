package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func respondErr(c *gin.Context, status int, msg string) {
	err := struct {
		Status int    `json:"status_code"`
		Error  string `json:"error_message"`
	}{
		status,
		msg,
	}

	c.JSON(status, err)
}

func respond(c *gin.Context, msg string) {
	res := struct {
		Msg string `json:"message"`
	}{msg}
	c.JSON(http.StatusOK, res)
}
