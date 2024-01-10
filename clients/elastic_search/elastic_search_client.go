package elastic_search

import (
	"context"
	"fmt"
	"github.com/lucaseg/go-bookstore-items-api/logger"
	"github.com/olivere/elastic"
	"time"
)

var (
	Client elasticClientInterface = &elasticClient{}
)

type elasticClientInterface interface {
	Index(string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string, string) (*elastic.GetResult, error)
	Search(string, elastic.Query) (*elastic.SearchResult, error)
}

type elasticClient struct {
	client *elastic.Client
}

func init() {
	esClient := elasticClient{}

	var err error
	esClient.client, err = elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		//elastic.SetErrorLog(),
		//elastic.SetInfoLog(),
		//elastic.SetHeaders()
	)
	if err != nil {
		panic(err)
	}
	Client = esClient
	logger.Info("Elastic search client is ready")
}

func (c elasticClient) Index(index string, document interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	idxResponse, err := c.client.Index().Index(index).Type("_doc").BodyJson(document).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("Error trying to index document in index : %s", index), err)
		return nil, err
	}

	return idxResponse, nil
}

func (c elasticClient) Get(index string, docType string, id string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := c.client.Get().Index(index).Type(docType).Id(id).Do(ctx)
	if err != nil {
		logger.Error("", err)
		return nil, err
	}
	if !result.Found {
		return nil, nil
	}
	return result, nil
}

func (c elasticClient) Search(index string, query elastic.Query) (*elastic.SearchResult, error) {
	ctx := context.Background()
	result, err := c.client.Search(index).Query(query).RestTotalHitsAsInt(true).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("Error trying to search documents for index : %s", index), err)
		return nil, err
	}
	return result, nil
}
