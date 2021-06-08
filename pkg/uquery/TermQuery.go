package uquery

import (
	"fmt"

	"github.com/blevesearch/bleve/v2"
	"github.com/prabhatsharma/mango/pkg/mbase"
)

func TermQuery(indexName string, iQuery mbase.MangoQuery) (*bleve.SearchResult, error) {
	requestedIndex := mbase.INDEX_LIST[indexName]

	dateQuery := bleve.NewDateRangeQuery(iQuery.Query.StartTime, iQuery.Query.EndTime)
	termQuery := bleve.NewTermQuery(iQuery.Query.Term)
	query := bleve.NewConjunctionQuery(dateQuery, termQuery)

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
