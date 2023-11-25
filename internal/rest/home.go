package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ "YUJIN.GG": "HOME" })
}
