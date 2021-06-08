package mindex

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/prabhatsharma/mango/pkg/mbase"
)

func InsertDocument(c *gin.Context) {
	if mbase.DATA_DIR == "" {
		mbase.DATA_DIR = "data"
	}

	indexName := c.Param("indexName")

	var doc map[string]interface{}

	requestedIndex := mbase.INDEX_LIST[indexName]

	err := c.BindJSON(&doc)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"result": "Bind error:",
			"error":  err.Error(),
		})
		return
	}

	doc["_timestamp_"] = time.Now()

	uuidWithHyphen := uuid.New()
	err = requestedIndex.Index(uuidWithHyphen.String(), doc)

	if err != nil {
		fmt.Println("error inserting document: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Indexing error:",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": doc,
	})
}
