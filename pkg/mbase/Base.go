package mbase

import "time"

type MangoQuery struct {
	SearchType string         `json:"searchType"`
	Size       int            `json:"size"`
	Explain    bool           `json:"explain"`
	Highlight  QueryHighlight `json:"highlight"`
	Fields     []string       `json:"fields"`
	Query      QueryParams    `json:"query"`
}

type QueryParams struct {
	Boost     int       `json:"boost"`
	Term      string    `json:"term"`
	Terms     []string  `json:"terms"` // For phrase query
	Field     string    `json:"field"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

type QueryHighlight struct {
	Fields []string `json:"fields"`
	Style  string   `json:"style"`
}
