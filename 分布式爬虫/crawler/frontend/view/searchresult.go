package view

import (
	"html/template"
	"io"
	"learn/crawler/frontend/model"
)

type SearchResultView struct {
	template *template.Template
}

func CreateSearchResultView(filename string) SearchResultView  {
	return SearchResultView{
		template:template.Must(
			template.ParseFiles(filename)),
	}
}
func (s SearchResultView) Render(w io.Writer, date model.SearchResult) error {
	return s.template.Execute(w, date)
}
