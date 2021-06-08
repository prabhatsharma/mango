package mindex

import (
	"fmt"
	"net/http"

	"github.com/blevesearch/bleve/v2"
	"github.com/gin-gonic/gin"
	"github.com/prabhatsharma/mango/pkg/mbase"
	"github.com/prabhatsharma/mango/pkg/uquery"
)

func SearchIndex(c *gin.Context) {
	if mbase.DATA_DIR == "" {
		mbase.DATA_DIR = "data"
	}

	indexName := c.Param("indexName")

	var iQuery mbase.MangoQuery

	c.BindJSON(&iQuery)

	var searchResult *bleve.SearchResult
	var err error

	switch iQuery.SearchType {
	case "term":
		searchResult, err = uquery.TermQuery(indexName, iQuery)
	case "querystring":
		searchResult, err = uquery.QueryStringQuery(indexName, iQuery)
	case "matchall":
		searchResult, err = uquery.MatchAllQuery(indexName, iQuery)
	case "match":
		searchResult, err = uquery.MatchQuery(indexName, iQuery)
	case "phrase":
		searchResult, err = uquery.PhraseQuery(indexName, iQuery)
	case "match phrase":
		searchResult, err = uquery.MatchPhraseQuery(indexName, iQuery)
	case "prefix":
		searchResult, err = uquery.PrefixQuery(indexName, iQuery)
	case "fuzzy":
		searchResult, err = uquery.FuzzyQuery(indexName, iQuery)
	case "doc id":
		searchResult, err = uquery.DocIDQuery(indexName, iQuery)
	case "date range":
		searchResult, err = uquery.DateRangeQuery(indexName, iQuery)

	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Query type needs to be one of: term, match, phrase, match phrase, prefix, fuzzy, conjunction, disjunction, boolean, numeric range, date range, query string, match all, match none, doc id ",
		})
		return
	}

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Could not open document in the index:",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": searchResult,
	})
}
