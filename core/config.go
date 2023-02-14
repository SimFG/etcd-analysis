package core

import "go.etcd.io/etcd/client/pkg/v3/transport"

var (
	C = Cfg{}
)

type Cfg struct {
	Endpoints []string
	TLS       transport.TLSInfo
}
