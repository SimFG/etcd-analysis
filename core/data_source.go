package core

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"math/rand"
	"os"
)

var (
	client *clientv3.Client
)

func InitClient() *clientv3.Client {
	connEndpoints := []string{"127.0.0.1:2379"}

	l := len(C.Endpoints)
	if l != 0 {
		connEndpoints = []string{C.Endpoints[rand.Intn(l)]}
	}

	cfg := clientv3.Config{
		Endpoints: connEndpoints,
	}

	c, err := clientv3.New(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dial error: %v\n", err)
		os.Exit(1)
	}
	client = c
	return client
}

func GetAllData() *clientv3.GetResponse {
	resp, err := client.Get(context.Background(), EmptyChar(), clientv3.WithRange(EmptyChar()),
		clientv3.WithSerializable(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	if err != nil {
		Exit(err)
	}
	return resp
}
