package client

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewProductClient, NewProduct_0Client)