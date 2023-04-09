package client

import (
	"context"
	"fmt"
	"gateway/pb"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// 返回多个GRPC的连接

func NewProduct_0Client(etcd *etcd.Registry, logger log.Logger) (pb.ProductClient, error) {

	connGRPC, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///product"),
		grpc.WithDiscovery(etcd),
	)
	if err != nil {
		//log.NewHelper(logger).WithContext(context.Background()).Errorw("kind", "grpc-client", "reason", "Grpc")
		//return nil, err
		fmt.Println(err.Error())
	}

	return pb.NewProductClient(connGRPC), nil
}
