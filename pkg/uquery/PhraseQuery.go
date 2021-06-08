package uquery

import (
	"fmt"

	"github.com/blevesearch/bleve/v2"
	"github.com/prabhatsharma/mango/pkg/mbase"
)

func PhraseQuery(indexName string, iQuery mbase.MangoQuery) (*bleve.SearchResult, error) {
	requestedIndex := mbase.INDEX_LIST[indexName]

	dateQuery := bleve.NewDateRangeQuery(iQuery.Query.StartTime, iQuery.Query.EndTime)
	matchQuery := bleve.NewPhraseQuery(iQuery.Query.Terms, iQuery.Query.Field)
	query := bleve.NewConjunctionQuery(dateQuery, matchQuery)

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
