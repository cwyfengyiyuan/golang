package client

import (
	"learn/crawler/engine"
	"learn/crawler_distributed/config"
	"learn/crawler_distributed/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	//client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkPort0))
	//if err != nil {
	//	return nil, err
	//}
	return func(req engine.Request) (result engine.ParseResult, e error) {
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult
		c := <- clientChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, nil
		}
		return worker.DeserializeResult(sResult), nil
	}
}
