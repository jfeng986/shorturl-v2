package logic

import (
	"context"

	"shorturl-v2/gateway/internal/svc"
	"shorturl-v2/gateway/internal/types"
	"shorturl-v2/rpc/transform/transformer"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortenLogic) Shorten(req *types.ShortenRequest) (*types.ShortenResponse, error) {
	resp, err := l.svcCtx.Transformer.Shorten(l.ctx, &transformer.ShortenRequest{
		OriginalURL: req.OriginalURL,
		CustomAlias: req.CustomAlias,
	})
	if err != nil {
		return nil, err
	}

	return &types.ShortenResponse{
		ShortURL: resp.ShortURL,
	}, nil
}
