// Code generated by goctl. DO NOT EDIT!
// Source: order.proto

package order

import (
	"context"



	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (

	Order interface {
		Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
		Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
		Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error)
		Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error)
		List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
		Paid(ctx context.Context, in *PaidRequest, opts ...grpc.CallOption) (*PaidResponse, error)
	}

	defaultOrder struct {
		cli zrpc.Client
	}
)

func NewOrder(cli zrpc.Client) Order {
	return &defaultOrder{
		cli: cli,
	}
}

func (m *defaultOrder) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	client := NewOrderClient(m.cli.Conn())
	return client.Create(ctx, in, opts...)
}

func (m *defaultOrder) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	client := NewOrderClient(m.cli.Conn())
	return client.Update(ctx, in, opts...)
}

func (m *defaultOrder) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	client := NewOrderClient(m.cli.Conn())
	return client.Remove(ctx, in, opts...)
}

func (m *defaultOrder) Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error) {
	client := NewOrderClient(m.cli.Conn())
	return client.Detail(ctx, in, opts...)
}

func (m *defaultOrder) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	client := NewOrderClient(m.cli.Conn())
	return client.List(ctx, in, opts...)
}

func (m *defaultOrder) Paid(ctx context.Context, in *PaidRequest, opts ...grpc.CallOption) (*PaidResponse, error) {
	client := NewOrderClient(m.cli.Conn())
	return client.Paid(ctx, in, opts...)
}
