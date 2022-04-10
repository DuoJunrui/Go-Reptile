package persist

import (
	"context"
	"github.com/olivere/elastic/v7"
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 1
		for {
			item := <-out
			log.Printf("Item Saver: get item #%d: %v", itemCount, item)
			itemCount++

			_, err := save(item)
			if err != nil {
				log.Printf("Item saver: error saving item %v: %v", item, err)
			}

		}
	}()
	return out
}

func save(item interface{}) (id string, err error) {
	// Must turn off sniff in docker
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return "", err
	}

	resp, err := client.Index().
		Index("userinfo").
		Type("vxiangqin").
		BodyJson(item).
		Do(context.Background())

	if err != nil {
		return "", err
	}

	return resp.Id, nil
}
