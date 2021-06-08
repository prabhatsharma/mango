package mindex

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDocument(c *gin.Context) {

	
	c.JSON(http.StatusOK, gin.H{
		"result": "DeleteDocument",
	})
}
