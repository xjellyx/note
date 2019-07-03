package etcd

import (
	"github.com/coreos/etcd/clientv3"
	"time"
)

// GetConfig 获取配置文件
func (c *Client) GetConfig() (ret *clientv3.Config) {
	ret = new(clientv3.Config)
	if c != nil {
		ret.TLS = c.TLS
		ret.Username = c.Username
		ret.Password = c.Password
		ret.DialTimeout = c.GetTimeout()
		ret.Endpoints = c.Endpoints
		ret.Context = c.Ctx

	}
	return
}

// GetTimeout 取超时
func (c *Client) GetTimeout() (ret time.Duration) {
	if c.Timeout > 0 {
		ret = time.Duration(c.Timeout) * time.Second
	}
	return
}
