package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"test-task/internal/models/elasticc"

	"github.com/elastic/go-elasticsearch/v8"
)

type SlotRepository struct {
	client *elasticsearch.Client
}

func NewSlotRepository(client *elasticsearch.Client) *SlotRepository {

	return &SlotRepository{
		client: client,
	}
}

func (r *SlotRepository) UpdateProductSlots(ctx context.Context, updates []elasticc.ProductSlotUpdate) error {

	var buf bytes.Buffer
	for _, item := range updates {

		meta := []byte(
			fmt.Sprintf(
				`{"update":{"_index":"products","_id":"%s"}}`,
				item.ID,
			),
		)

		buf.Write(meta)
		buf.WriteByte('\n')

		doc := map[string]interface{}{
			"doc": map[string]interface{}{
				"slot": item.Slot,
			},
		}

		data, err := json.Marshal(doc)

		if err != nil {
			return err
		}

		buf.Write(data)
		buf.WriteByte('\n')
	}

	res, err := r.client.Bulk(
		bytes.NewReader(buf.Bytes()),
		r.client.Bulk.WithContext(ctx),
	)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf(
			"elastic bulk update error: %s",
			res.String(),
		)
	}

	return nil
}
