package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Health to manage server health
type Health struct{}

//Check show the status of server
func (h Health) Check(c *gin.Context) {
	c.String(http.StatusOK, "Running")
}
