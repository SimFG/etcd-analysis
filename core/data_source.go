package core

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"math/rand"
	"os"
)

const (
	readCount = 1000
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

func GetAllData() (*clientv3.GetResponse, <-chan []*mvccpb.KeyValue) {
	c := make(chan []*mvccpb.KeyValue, 10)

	getFunc := func(start string, count int64) *clientv3.GetResponse {
		resp, err := client.Get(context.Background(), start, clientv3.WithFromKey(),
			clientv3.WithSerializable(),
			//clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend),
			clientv3.WithLimit(count))
		if err != nil {
			Exit(err)
		}
		return resp
	}
	resp := getFunc(EmptyChar(), 2)
	c <- resp.Kvs

	go func(first *clientv3.GetResponse) {
		defer close(c)
		nextResp := first
		l := len(nextResp.Kvs)
		if l <= 1 {
			return
		}
		for {
			lastKey := nextResp.Kvs[l-1].Key
			nextResp = getFunc(string(lastKey), readCount)
			l = len(nextResp.Kvs)
			if l <= 1 {
				return
			}
			c <- nextResp.Kvs[1:]
		}
	}(resp)

	return resp, c
}
