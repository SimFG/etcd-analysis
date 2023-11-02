package core

import (
	"fmt"
	"math/rand"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	readCount = 1000
)

var client *clientv3.Client

func InitClient() *clientv3.Client {
	connEndpoints := []string{"127.0.0.1:2379"}

	l := len(C.Endpoints)
	if l != 0 {
		connEndpoints = []string{C.Endpoints[rand.Intn(l)]}
	}

	cfg := clientv3.Config{
		Endpoints: connEndpoints,
	}

	if C.TLS.CertFile != "" && C.TLS.KeyFile != "" && C.TLS.TrustedCAFile != "" {
		tlsConfig, err := C.TLS.ClientConfig()
		if err != nil {
			Exit(fmt.Errorf("failed to get etcd tls config, err is %v", err))
		}
		cfg.TLS = tlsConfig
		cfg.TLS.InsecureSkipVerify = true
	}

	c, err := clientv3.New(cfg)
	if err != nil {
		Exit(fmt.Errorf("dial error: %v\n", err))
	}
	EtcdOpTimeout = time.Duration(C.CommandTimeout) * time.Second
	err = EtcdStatus(c)
	if err != nil {
		Exit(fmt.Errorf("unvaliable etcd server, error: %v\n", err))
	}
	client = c
	return client
}

func GetDataWithPrefix(prefix string) (*clientv3.GetResponse, <-chan []*mvccpb.KeyValue) {
	c := make(chan []*mvccpb.KeyValue, 10)

	getFunc := func(start string, count int64) *clientv3.GetResponse {
		opts := []clientv3.OpOption{
			clientv3.WithSerializable(),
			clientv3.WithLimit(count),
		}
		if prefix != "" {
			opts = append(opts, clientv3.WithPrefix())
		} else {
			opts = append(opts, clientv3.WithFromKey())
		}
		resp, err := EtcdGet(client, start, opts...)
		if err != nil {
			Exit(err)
		}
		return resp
	}
	if prefix != "" {
		resp := getFunc(prefix, 0)
		c <- resp.Kvs
		close(c)
		return resp, c
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

func GetAllData() (*clientv3.GetResponse, <-chan []*mvccpb.KeyValue) {
	return GetDataWithPrefix("")
}
