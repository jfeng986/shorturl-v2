package logic

import (
	"context"

	"shorturl-v2/rpc/transform/internal/svc"
	"shorturl-v2/rpc/transform/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShortURLLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShortURLLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShortURLLogic {
	return &GetShortURLLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetShortURLLogic) GetShortURL(in *transform.ShortenRequest) (*transform.ShortenResponse, error) {
	// todo: add your logic here and delete this line

	return &transform.ShortenResponse{}, nil
}
