package third_party

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"inventory-management/backend/internal/model"
	"strconv"
)

type UserElasticsearch struct {
	Elasticsearch *elasticsearch.Client
}

func NewUserElasticsearch(elasticsearch *elasticsearch.Client) UserElasticsearchContract {
	return &UserElasticsearch{
		Elasticsearch: elasticsearch,
	}
}

func (service *UserElasticsearch) CountAll(ctx context.Context, data bytes.Buffer) (int64, error) {
	count, err := service.Elasticsearch.Count(
		service.Elasticsearch.Count.WithContext(ctx),
		service.Elasticsearch.Count.WithIndex("users"),
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

func (service *UserElasticsearch) Search(ctx context.Context, data bytes.Buffer, offset int, limit int) (map[string]interface{}, error) {
	res, err := service.Elasticsearch.Search(
		service.Elasticsearch.Search.WithContext(ctx),
		service.Elasticsearch.Search.WithIndex("users"),
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

func (service *UserElasticsearch) Create(ctx context.Context, user *model.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	reqEs := esapi.IndexRequest{
		Index:      "users",
		DocumentID: strconv.FormatInt(user.ID, 10),
		Body:       bytes.NewReader(data),
	}

	res, err := reqEs.Do(ctx, service.Elasticsearch)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func (service *UserElasticsearch) Delete(ctx context.Context, id int64) error {
	idString := strconv.FormatInt(id, 10)
	res, err := service.Elasticsearch.Delete("users", idString,
		service.Elasticsearch.Delete.WithContext(ctx),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
