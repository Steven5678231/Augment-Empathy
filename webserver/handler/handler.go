package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//DefaultHomePageHandler:
func DefaultHomePageHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/index.html")
}
