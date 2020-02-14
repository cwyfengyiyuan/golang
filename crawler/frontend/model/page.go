package model

type SearchResult struct {
	Hits int64
	Start int
	PrevFrom int
	NextFrom int
	Query string
	Items []interface{}
}
