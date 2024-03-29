package main

import (
	"flag"
	"fmt"
	"leango/crawler_distributed/config"
	"leango/crawler_distributed/persist"
	"leango/crawler_distributed/rpcsupport"
	"log"

	"github.com/olivere/elastic/v7"
)

var port = flag.Int("port", 0, "the port for me to listen to")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.SaverRpc(host, &persist.ItemSaveService{Client: client, Index: index})
}
