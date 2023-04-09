package client

import (
	"context"
	"fmt"
	"gateway/pb/category/v1"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"time"
)

// 返回多个GRPC的连接

func NewProductClient(etcd *etcd.Registry, logger log.Logger) (v1.CategoryClient, error) {

	connGRPC, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///product"),
		grpc.WithDiscovery(etcd),
		grpc.WithTimeout(time.Second*5), // 连接超时时间
	)
	if err != nil {
		//log.NewHelper(logger).WithContext(context.Background()).Errorw("kind", "grpc-client", "reason", "Grpc")
		//return nil, err
		fmt.Println(err.Error())

	}

	return v1.NewCategoryClient(connGRPC), nil
}
