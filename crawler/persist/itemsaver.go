package persist

import (
	"context"
	"errors"
	"gopkg.in/olivere/elastic.v5"
	"learn/crawler/engine"
	"log"
)
func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://192.168.200.17:9200"))
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Saver: got item " + "#%d: %v", itemCount, item)
			itemCount++
			err := save(client, index, item)
			if err != nil {
				log.Printf("Item Saver: error"+"saving item %v: %v", item, err)
			}

		}
	}()
	return out, nil
}
//func ItemSaver() chan interface{} {
//	out := make(chan interface{})
//	go func() {
//		itemCount := 0
//		for {
//			item := <- out
//			log.Printf("Item Saver: got item " + "#%d: %v", itemCount, item)
//			itemCount++
//			//_, err := save(item)
//			//if err != nil {
//			//	log.Printf("Item Saver: error"+"saving item %v: %v", item, err)
//			//}
//
//		}
//	}()
//	return out
//}
func save(client *elastic.Client, index string, item engine.Item) error {
	if item.Type == "" {
		return errors.New("must supply Type")
	}
	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
