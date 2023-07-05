package third_party

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v7"
	"reflect"
	"strconv"
)

type Elasticsearch struct {
	Elasticsearch *elasticsearch.Client
}

func NewElasticsearch(elasticsearch *elasticsearch.Client) ElasticsearchContract {
	return &Elasticsearch{
		Elasticsearch: elasticsearch,
	}
}

func (service *Elasticsearch) CountAll(ctx context.Context, index string, data bytes.Buffer) (int64, error) {
	count, err := service.Elasticsearch.Count(
		service.Elasticsearch.Count.WithContext(ctx),
		service.Elasticsearch.Count.WithIndex(index),
		service.Elasticsearch.Count.WithBody(&data),
		service.Elasticsearch.Count.WithPretty(),
	)
	if err != nil {
		return 0, err
	}
	defer count.Body.Close()

	var totalHits map[string]interface{}
	if err := json.NewDecoder(count.Body).Decode(&totalHits); err != nil {
		return 0, err
	}

	var totalHitsInt int64
	if totalHits["count"] != nil {
		totalHitsInt = int64(totalHits["count"].(float64))
	}

	return totalHitsInt, nil
}

func (service *Elasticsearch) Search(ctx context.Context, index string, data bytes.Buffer, offset int, limit int) (map[string]interface{}, error) {
	res, err := service.Elasticsearch.Search(
		service.Elasticsearch.Search.WithContext(ctx),
		service.Elasticsearch.Search.WithIndex(index),
		service.Elasticsearch.Search.WithBody(&data),
		service.Elasticsearch.Search.WithTrackTotalHits(true),
		service.Elasticsearch.Search.WithPretty(),
		service.Elasticsearch.Search.WithSize(limit),
		service.Elasticsearch.Search.WithFrom(offset),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var searchResponse map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&searchResponse); err != nil {
		return nil, err
	}

	return searchResponse, nil
}

func (service *Elasticsearch) Create(ctx context.Context, index string, request interface{}, id int64) error {
	if reflect.TypeOf(request).Kind() != reflect.Struct {
		return errors.New("the request is not a struct type")
	}

	data, err := json.Marshal(request)
	if err != nil {
		return err
	}

	idString := strconv.FormatInt(id, 10)
	dataReader := bytes.NewReader(data)
	res, err := service.Elasticsearch.Create(
		index,
		idString,
		dataReader,
		service.Elasticsearch.Create.WithContext(ctx),
		service.Elasticsearch.Create.WithPretty(),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func (service *Elasticsearch) Update(ctx context.Context, index string, request interface{}, id int64) error {
	if reflect.TypeOf(request).Kind() != reflect.Struct {
		return errors.New("the request is not a struct type")
	}

	err := service.Delete(ctx, index, id)
	if err != nil {
		return err
	}

	err = service.Create(ctx, index, request, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *Elasticsearch) Delete(ctx context.Context, index string, id int64) error {
	idString := strconv.FormatInt(id, 10)
	res, err := service.Elasticsearch.Delete(
		index,
		idString,
		service.Elasticsearch.Delete.WithContext(ctx),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
