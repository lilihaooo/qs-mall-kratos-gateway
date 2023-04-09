package registry

import (
	"gateway/internal/conf"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	etcdclient "go.etcd.io/etcd/client/v3"
	"time"
)

func NewEtcdRegistry(conf *conf.Etcd) *etcd.Registry {
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints:   []string{conf.Url},
		DialTimeout: 5 * time.Second, // 连接超时时间
	})
	if err != nil {
		panic(err)
	}
	return etcd.New(client)
}
