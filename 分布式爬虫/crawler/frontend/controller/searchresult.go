package controller

import (
	"context"
	"gopkg.in/olivere/elastic.v5"
	"learn/crawler/engine"
	"learn/crawler/frontend/model"
	"learn/crawler/frontend/view"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler  {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://192.168.200.17:9200"))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view: view.CreateSearchResultView(template),
		client: client,
	}
}

//localhost:8888/search?q=男&from=20
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}
	page, err := h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h SearchResultHandler)  getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	resp, err := h.client.Search("zhenai_date").
		Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).
		From(from).
		Do(context.Background())
	if err != nil {
		return result, err
	}
	result.Query = q
	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)
	return result, nil
}
func rewriteQueryString(q string) string  {
	re := regexp.MustCompile(`([A-z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}
////取所有
//res, err = client.Search("megacorp").Type("employee").Do(context.Background())
//printEmployee(res, err)
//
////字段相等
//q := elastic.NewQueryStringQuery("last_name:Smith")
//res, err = client.Search("megacorp").Type("employee").Query(q).Do(context.Background())
//if err != nil {
//println(err.Error())
//}
//条件查询
//年龄大于30岁的
//boolQ := elastic.NewBoolQuery()
//boolQ.Must(elastic.NewMatchQuery("last_name", "smith"))
//boolQ.Filter(elastic.NewRangeQuery("age").Gt(30))
//res, err = client.Search("megacorp").Type("employee").Query(q).Do(context.Background())
//printEmployee(res, err)
//
////短语搜索 搜索about字段中有 rock climbing
//matchPhraseQuery := elastic.NewMatchPhraseQuery("about", "rock climbing")
//res, err = client.Search("megacorp").Type("employee").Query(matchPhraseQuery).Do(context.Background())
//printEmployee(res, err)
//
////分析 interests
//aggs := elastic.NewTermsAggregation().Field("interests")
//res, err = client.Search("megacorp").Type("employee").Aggregation("all_interests", aggs).Do(context.Background())
//printEmployee(res, err)
//简单分页
//func list(size,page int) {
//	if size < 0 || page < 1 {
//		fmt.Printf("param error")
//		return
//	}
//	res,err := client.Search("megacorp").
//		Type("employee").
//		Size(size).
//		From((page-1)*size).
//		Do(context.Background())
//	printEmployee(res, err)
//
//}

