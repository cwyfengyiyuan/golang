package view

import (
	"learn/crawler/engine"
	"learn/crawler/frontend/model"
	common "learn/crawler/model"
	"os"
	"testing"
)

func TestTemplate(t *testing.T)  {
	view := CreateSearchResultView("template.html")
	out, err := os.Create("template_test.html")
	page := model.SearchResult{}
	page.Hits = 5
	item := engine.Item{
		Url:"https://album.zhenai.com/u/101414547",
		Type:"zhenai",
		Id:"101414547",
		Payload:common.Profile {
			Name:"Stellapoo",
			Gender:"女士",
			User: []string{"name:xiuse","离异","23岁","有车","有房"},
		},
	}
	for i := 0; i < 5; i++ {
		page.Items = append(page.Items, item)
	}
	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}