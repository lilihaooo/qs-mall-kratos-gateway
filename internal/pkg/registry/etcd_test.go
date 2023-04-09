package registry

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	etcdclient "go.etcd.io/etcd/client/v3"
	"testing"
)

func Test_Etcd(t *testing.T) {
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}
	cli := etcd.New(client)
	err = cli.Register(context.Background(), &registry.ServiceInstance{
		ID:        "user-service",
		Name:      "product",
		Version:   "v1.0.0",
		Endpoints: []string{"grpc://127.0.0.1:9000"},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("创建服务成功")
}
