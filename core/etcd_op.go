package core

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var EtcdOpTimeout time.Duration

func EtcdPut(etcdCli *clientv3.Client, key, val string, opts ...clientv3.OpOption) error {
	ctx, cancel := context.WithTimeout(context.Background(), EtcdOpTimeout)
	defer cancel()
	_, err := etcdCli.Put(ctx, key, val, opts...)
	return err
}

func EtcdGet(etcdCli *clientv3.Client, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), EtcdOpTimeout)
	defer cancel()
	resp, err := etcdCli.Get(ctx, key, opts...)
	return resp, err
}

func EtcdDelete(etcdCli *clientv3.Client, key string, opts ...clientv3.OpOption) error {
	ctx, cancel := context.WithTimeout(context.Background(), EtcdOpTimeout)
	defer cancel()
	_, err := etcdCli.Delete(ctx, key, opts...)
	return err
}

func EtcdTxn(etcdCli *clientv3.Client, fun func(txn clientv3.Txn)) {
	ctx, cancel := context.WithTimeout(context.Background(), EtcdOpTimeout)
	defer cancel()
	etcdTxn := etcdCli.Txn(ctx)
	fun(etcdTxn)
}

func EtcdStatus(etcdCli *clientv3.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), EtcdOpTimeout)
	defer cancel()
	for _, endpoint := range etcdCli.Endpoints() {
		_, err := etcdCli.Status(ctx, endpoint)
		if err != nil {
			return err
		}
	}
	return nil
}
