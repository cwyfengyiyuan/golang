package main

import (
	"learn/crawler/engine"
	"learn/crawler/persist"
	"learn/crawler/scheduler"
	"learn/crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("zhenai_date")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 20,
		ItemChan: itemChan,
	}
	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
	//e.Run(engine.Request{
	//	Url:       "http://www.zhenai.com/zhenghun/shanghai",
	//	ParseFunc: parser.ParseCity,
	//})
}