package svc

import (
	"github.com/fans1992/lebron/apps/app/api/internal/config"
	"github.com/fans1992/lebron/apps/order/rpc/order"
	"github.com/fans1992/lebron/apps/product/rpc/product"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	OrderRPC   order.Order
	ProductRPC product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRPC:   order.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
		ProductRPC: product.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
	}
}
