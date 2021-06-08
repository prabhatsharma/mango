# Mango

Lightweight modern search tool written in Go


API endpoints

PUT /api/:indexName - create new index
GET /api/:indexName - get index details
DELETE /api/:indexName - delete index
GET /api/_list - list indexes

PUT /api/:indexName/:docID - index document
GET /api/:indexName/_count - count documents in index
GET /api/:indexName/:docID - return stored fields of document
DELETE /api/:indexName/:docID - delete document


POST /api/:indexName/_search - search index
GET /api/:indexName/_fields - list fields used by documents in index
GET /api/:indexName/stats - list fields used by documents in index
GET /api/:indexName/:docID/_debug - return rows in index related to document


Search syntax

POST /api/:indexName/_search

<pre>
{
    "size": 10, # not implemented 
    "searchType": "match", # implemented  - match, phrase, match phrase, prefix, fuzzy
    # term, match, phrase, match phrase, prefix, fuzzy, conjunction, disjunction, boolean, numeric range, date range, query string, match all, match none, doc id 
    "from": 0, # not implemented 
    "explain": true, # not implemented 
    "highlight": {}, # not implemented 
    "query": {
        "boost": 1, # not implemented 
        "term": "family"
    },
    "fields": [
        "*"
    ]
}
</pre>

