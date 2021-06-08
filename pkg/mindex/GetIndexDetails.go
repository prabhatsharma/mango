package mindex

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIndexDetails(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "GetIndexDetails",
	})
}
