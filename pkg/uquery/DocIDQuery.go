package uquery

import (
	"fmt"

	"github.com/blevesearch/bleve/v2"
	"github.com/prabhatsharma/mango/pkg/mbase"
)

func DocIDQuery(indexName string, iQuery mbase.MangoQuery) (*bleve.SearchResult, error) {
	requestedIndex := mbase.INDEX_LIST[indexName]

	query := bleve.NewDocIDQuery(iQuery.Query.Terms)
	searchRequest := bleve.NewSearchRequest(query)
	searchRequest.Highlight = bleve.NewHighlight()
	searchRequest.Fields = iQuery.Fields

	searchResult, err := requestedIndex.Search(searchRequest)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return searchResult, nil
}
