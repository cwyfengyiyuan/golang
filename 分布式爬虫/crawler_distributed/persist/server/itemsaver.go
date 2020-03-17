package main

import (
	"flag"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"learn/crawler_distributed/config"
	"learn/crawler_distributed/persist"
	"learn/crawler_distributed/rpcsupport"
	"log"
)
var port = flag.Int("port", 0, "the port for me to listen on")
func main()  {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
	//log.Fatal(serveRpc(fmt.Sprintf(":%d", config.ItemSavePort), config.ElasticIndex))
}
func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://192.168.200.17:9200"))
	if err != nil {
		return err
	}
	return  rpcsupport.ServeRpc(host, &persist.ItemSaverService{
				Client: client,
				Index: index,
			})
}