// Code generated by iocli, run 'iocli gen' to re-generate

package api

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	dto "github.com/alibaba/ioc-golang/example/autowire/autowire_rpc/server/pkg/dto"
	rpc_client "github.com/alibaba/ioc-golang/extension/autowire/rpc/rpc_client"
)

func init() {
	rpc_client.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &serviceStructIOCRPCClient{}
		},
	})
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &serviceStructIOCRPCClient_{}
		},
	})
}

type serviceStructIOCRPCClient_ struct {
	GetUser_ func(name string, age int) (*dto.User, error)
}

func (s *serviceStructIOCRPCClient_) GetUser(name string, age int) (*dto.User, error) {
	return s.GetUser_(name, age)
}

type ServiceStructIOCRPCClient interface {
	GetUser(name string, age int) (*dto.User, error)
}

type serviceStructIOCRPCClient struct {
	GetUser func(name string, age int) (*dto.User, error)
}
