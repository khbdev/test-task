package elastic

import (
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

func BootstrapIndex(es *elasticsearch.Client) error {

	exists, err := es.Indices.Exists([]string{"products"})
	if err != nil {
		return err
	}

	if exists.StatusCode == 200 {
		fmt.Println("products index already exists")
		return nil
	}

	body := strings.NewReader(`
	{
	  "mappings": {
	    "properties": {
	      "id": {
	        "type": "keyword"
	      },
	      "company_id": {
	        "type": "keyword"
	      },
	      "name": {
	        "type": "text"
	      },
	      "sku": {
	        "type": "keyword"
	      },
	      "barcode": {
	        "type": "keyword"
	      },
	      "retail_price": {
	        "type": "double"
	      },
	      "slot": {
	        "type": "keyword"
	      }
	    }
	  }
	}
	`)

	resp, err := es.Indices.Create(
		"products",
		es.Indices.Create.WithBody(body),
	)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.IsError() {
		return fmt.Errorf("create index error: %s", resp.String())
	}

	fmt.Println("products index created")

	return nil
}
