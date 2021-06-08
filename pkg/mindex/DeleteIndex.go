package mindex

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prabhatsharma/mango/pkg/mbase"
)

func DeleteIndex(c *gin.Context) {
	if mbase.DATA_DIR == "" {
		mbase.DATA_DIR = "data"
	}

	indexName := c.Param("indexName")
	requestedIndex := mbase.INDEX_LIST[indexName]

	requestedIndex.Close()
	delete(mbase.INDEX_LIST, indexName)

	err := os.RemoveAll(mbase.DATA_DIR + "/" + indexName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result": "Could not delete Index: " + indexName,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Deleted Index: " + indexName,
	})

}
