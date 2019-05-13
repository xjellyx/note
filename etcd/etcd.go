package etcd

import (
	"context"
	"crypto/tls"
	"github.com/coreos/etcd/clientv3"
	"sync"
)

type Client struct {
	closeChan chan struct{}    // 关闭通道
	client    *clientv3.Client // v3 client
	leaseID   clientv3.LeaseID
	wg        sync.WaitGroup
	Endpoints []string `json:"endpoints"`
	// Username is a user name for authentication.
	Username string `json:"username"`

	// Password is a password for authentication.
	Password string `json:"password"`

	// 证书
	PathCert string `json:"patchCert"`
	PathKey  string `json:"patchKey"`
	PathCa   string `json:"patchCa"`

	// 超时
	Timeout int `json:"timeout"`

	// TLS holds the client secure credentials, if any.
	TLS *tls.Config `json:"-" yaml:"-"`

	//
	Ctx context.Context
}

func NewClient(s *Client) (ret *Client, err error) {
	if s != nil {
		ret = s
	} else {
		ret = new(Client)
	}

	if ret.TLS == nil && len(ret.PathCert) > 0 {
		//TODO
	}

	if ret.Timeout == 0 {
		ret.Timeout = 5
	}

	return
}
