// Code generated by goctl. DO NOT EDIT.
// Source: transform.proto

package transformer

import (
	"context"

	"shorturl-v2/rpc/transform/transform"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Empty                  = transform.Empty
	GetOriginalURLRequest  = transform.GetOriginalURLRequest
	GetOriginalURLResponse = transform.GetOriginalURLResponse
	ShortenRequest         = transform.ShortenRequest
	ShortenResponse        = transform.ShortenResponse

	Transformer interface {
		GetShortURL(ctx context.Context, in *ShortenRequest, opts ...grpc.CallOption) (*ShortenResponse, error)
		GetOriginalURL(ctx context.Context, in *GetOriginalURLRequest, opts ...grpc.CallOption) (*GetOriginalURLResponse, error)
		Redirect(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	}

	defaultTransformer struct {
		cli zrpc.Client
	}
)

func NewTransformer(cli zrpc.Client) Transformer {
	return &defaultTransformer{
		cli: cli,
	}
}

func (m *defaultTransformer) GetShortURL(ctx context.Context, in *ShortenRequest, opts ...grpc.CallOption) (*ShortenResponse, error) {
	client := transform.NewTransformerClient(m.cli.Conn())
	return client.GetShortURL(ctx, in, opts...)
}

func (m *defaultTransformer) GetOriginalURL(ctx context.Context, in *GetOriginalURLRequest, opts ...grpc.CallOption) (*GetOriginalURLResponse, error) {
	client := transform.NewTransformerClient(m.cli.Conn())
	return client.GetOriginalURL(ctx, in, opts...)
}

func (m *defaultTransformer) Redirect(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	client := transform.NewTransformerClient(m.cli.Conn())
	return client.Redirect(ctx, in, opts...)
}
