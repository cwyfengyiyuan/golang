package main

import (
	"learn/crawler/engine"
	"learn/crawler/model"
	"learn/crawler_distributed/config"
	"learn/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSave(t *testing.T)  {
	const host  = ":1234"
	//start ItemSaverServer
	go serveRpc(host, "test1")
	time.Sleep(time.Second*6)

	//start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	//call save
	item := engine.Item{
		Url:"https://album.zhenai.com/u/101414547",
		Type:"zhenai",
		Id:"101414547",
		Payload:model.Profile {
			Name:"Atismiss",
			Gender:"女士",
			User: []string{"离异","23岁","有车","有房"},
		},
	}
	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s;", result, err)
	}
}
