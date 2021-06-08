package mindex

import (
	"net/http"

	"github.com/blevesearch/bleve/v2"
	"github.com/gin-gonic/gin"
	"github.com/prabhatsharma/mango/pkg/mbase"
)

type MangoIndex struct {
	Name          string           `json:"name"`
	DocumentCount uint64           `json:"docCount"`
	Fields        []string         `json:"fields"`
	Stats         *bleve.IndexStat `json:"stats"`
}

func ListIndexes(c *gin.Context) {
	var ivL []MangoIndex
	var iv MangoIndex

	for k, v := range mbase.INDEX_LIST {
		iv.Name = k
		iv.DocumentCount, _ = v.DocCount()
		iv.Fields, _ = v.Fields()
		iv.Stats = v.Stats()

		ivL = append(ivL, iv)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": ivL,
	})
}
