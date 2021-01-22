package etcd

import (
	"context"
	"crypto/tls"
	"github.com/coreos/etcd/clientv3"
	"gorm.io/gorm"
	"sync"
)

// Client 客户端
type Client struct {
	gorm.Model
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

// Client 客户端
type Client２ struct {
	gorm.Model
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
