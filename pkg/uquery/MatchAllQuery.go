package uquery

import (
	"fmt"

	"github.com/blevesearch/bleve/v2"
	"github.com/prabhatsharma/mango/pkg/mbase"
)

func MatchAllQuery(indexName string, iQuery mbase.MangoQuery) (*bleve.SearchResult, error) {
	requestedIndex := mbase.INDEX_LIST[indexName]

	matchAllQuery := bleve.NewMatchAllQuery()
	dateQuery := bleve.NewDateRangeQuery(iQuery.Query.StartTime, iQuery.Query.EndTime)
	query := bleve.NewConjunctionQuery(dateQuery, matchAllQuery)

	searchRequest := bleve.NewSearchRequest(query)
	searchRequest.Highlight = bleve.NewHighlight()
	searchRequest.Fields = iQuery.Fields
	searchRequest.Size = iQuery.Size

	searchResult, err := requestedIndex.Search(searchRequest)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return searchResult, nil
}
