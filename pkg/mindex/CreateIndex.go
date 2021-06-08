package mindex

import (
	"fmt"
	"net/http"

	"github.com/blevesearch/bleve/v2"
	"github.com/gin-gonic/gin"
	"github.com/prabhatsharma/mango/pkg/mbase"
)

// CreateIndex creates a new Index
func CreateIndex(c *gin.Context) {
	if mbase.DATA_DIR == "" {
		mbase.DATA_DIR = "data"
	}

	indexName := c.Param("indexName")

	// open a new index
	mapping := bleve.NewIndexMapping()
	timestampMapping := bleve.NewDateTimeFieldMapping()
	mapping.DefaultMapping.AddFieldMappingsAt("_timestamp_", timestampMapping)
	index, err := bleve.New(mbase.DATA_DIR+"/"+indexName, mapping)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Error occured creating index: " + indexName,
			"error":  err.Error(),
		})
		return
	}

	mbase.INDEX_LIST[indexName] = index

	c.JSON(http.StatusCreated, gin.H{
		"result": "Index " + index.Name() + " created. ",
	})
}
