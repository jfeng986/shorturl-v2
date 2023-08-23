package logic

import (
	"context"

	"shorturl-v2/rpc/transform/internal/svc"
	"shorturl-v2/rpc/transform/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedirectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRedirectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedirectLogic {
	return &RedirectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RedirectLogic) Redirect(in *transform.Empty) (*transform.Empty, error) {
	// todo: add your logic here and delete this line

	return &transform.Empty{}, nil
}
