package persist

import (
	"context"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"learn/crawler/engine"
	"learn/crawler/model"
	"testing"
)

func TestSaver(t *testing.T) {
	exp := engine.Item{
		Url:"https://album.zhenai.com/u/101414547",
		Type:"zhenai",
		Id:"101414547",
		Payload:model.Profile {
			Name:"Stellapoo",
			Gender:"女士",
			User: []string{"name:xiuse","离异","23岁","有车","有房"},
		},
	}
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://192.168.200.17:9200"))
	if err != nil {
		panic(err)
	}
	const index  = "test_zhenai"
	err = save(client, index, exp)
	if err != nil {
		panic(err)
	}
	resp, err := client.Get().
		Index(index).
		Type(exp.Type).
		Id(exp.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", resp.Source)

	var actual engine.Item
	json.Unmarshal(*resp.Source, &actual)

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	//if actual != exp {
	//	t.Errorf("got %v; expected %v", actual, exp)
	//}else {
	//	t.Logf("Yes, %+v, %+v", actual, exp)
	//}
	t.Logf("%+v", actual)
	t.Logf("%+v", exp)
}
