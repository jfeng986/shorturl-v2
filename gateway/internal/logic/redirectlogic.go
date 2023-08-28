package logic

import (
	"context"

	"shorturl-v2/gateway/internal/svc"
	"shorturl-v2/gateway/internal/types"
	"shorturl-v2/rpc/transform/transformer"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedirectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRedirectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedirectLogic {
	return &RedirectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RedirectLogic) Redirect(req *types.RedirectRequest) (*types.RedirectResponse, error) {
	resp, err := l.svcCtx.Transformer.Expand(l.ctx, &transformer.ExpandRequest{
		ShortURL: req.ShortURL,
	})
	if err != nil {
		return nil, err
	}
	return &types.RedirectResponse{
		OriginalURL: resp.OriginalURL,
	}, nil
}
