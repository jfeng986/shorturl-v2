package logic

import (
	"context"

	"shorturl-v2/gateway/internal/svc"
	"shorturl-v2/gateway/internal/types"
	"shorturl-v2/rpc/transform/transformer"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExpandLogic) Expand(req *types.ExpandRequest) (*types.ExpandResponse, error) {
	resp, err := l.svcCtx.Transformer.Expand(l.ctx, &transformer.ExpandRequest{
		ShortURL: req.ShortURL,
	})
	if err != nil {
		return nil, err
	}
	return &types.ExpandResponse{
		OriginalURL: resp.OriginalURL,
	}, nil
}
