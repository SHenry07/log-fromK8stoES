package es

import (
	"context"
	"log"
	"strings"

	"github.com/olivere/elastic/v7"
)

type LogEntry struct {
	Topic string
	Timestamp string
	Data string
}

var (
	client *elastic.Client
	Datachan chan *LogEntry

)
// Init build a es's client
func Init(urls []string, ChanMaxSize int ) {
	for i := 0; i < len(urls); i++ {
		if !strings.HasPrefix(urls[i],"http://") {
			urls[i] = "http://"+urls[i]
		}
	}
	var err error
	client, err = elastic.NewClient(elastic.SetURL(urls...))
	if err != nil {
		// Handle error
		log.Panicf("[ElasticSearch]: Can't establish to ES %v", err)
	}
	log.Println("[ElasticSearch]: connect success")
	Datachan = make(chan *LogEntry, ChanMaxSize)
}

// SendToEs ....
func SendToEs(index string) {
	for {
		LogEntry := <-Datachan
		put1, err := client.Index().
			Index(index).
			BodyJson(&LogEntry).
			Do(context.Background())
		if err != nil {
			// Handle error
			panic(err)
		}
		log.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	}
}
