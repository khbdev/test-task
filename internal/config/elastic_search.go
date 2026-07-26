package config

import (
	"fmt"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

func NewElasticsearchClient() *elasticsearch.Client {
	esURL := os.Getenv("ELASTICSEARCH_URL")
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{esURL},
	})
	if err != nil {
		log.Fatalf("failed to create Elasticsearch client: %v", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("failed to connect to Elasticsearch: %v", err)
	}
	defer res.Body.Close()
	fmt.Println("Successfully connected to Elasticsearch")

	return es
}
