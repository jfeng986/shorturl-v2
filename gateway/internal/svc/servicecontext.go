package svc

import (
	"shorturl-v2/gateway/internal/config"
	"shorturl-v2/rpc/qrcode/qrcoder"
	"shorturl-v2/rpc/transform/transformer"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	Transformer transformer.Transformer
	Qrcoder     qrcoder.Qrcoder
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		Transformer: transformer.NewTransformer(zrpc.MustNewClient(c.Transform)),
		Qrcoder:     qrcoder.NewQrcoder(zrpc.MustNewClient(c.Qrcode)),
	}
}
