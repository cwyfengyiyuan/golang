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
		WorkerCount: 3,
		ItemChan: itemChan,
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})
	//e.Run(engine.Request{
	//	Url:       "http://www.zhenai.com/zhenghun/shanghai",
	//	ParseFunc: parser.ParseCity,
	//})
}