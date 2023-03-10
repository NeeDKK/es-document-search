package config

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

var ctx = context.Background()
var EsClient *elastic.Client

func InitEs() {
	client, err := elastic.NewClient(
		elastic.SetSniff(false), elastic.SetURL(GlobalConfig.Elasticsearch.Url),
	)
	if err != nil {
		log.Fatal("es 连接失败:", err)
	}
	// ping通服务端，并获得服务端的es版本
	info, code, err := client.Ping(GlobalConfig.Elasticsearch.Url).Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Elasticsearch call code:", code, " version:", info.Version.Number)
	EsClient = client
}
