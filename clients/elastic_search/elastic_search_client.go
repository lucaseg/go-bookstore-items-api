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
	idxResponse, err := c.client.Index().Index(index).Type("item").BodyJson(document).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("Error trying to index document in index : %s", index), err)
		return nil, err
	}

	return idxResponse, nil
}
