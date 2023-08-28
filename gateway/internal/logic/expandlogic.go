package logic

import (
	"context"

	"shorturl-v2/gateway/internal/svc"
	"shorturl-v2/gateway/internal/types"

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

func (l *ExpandLogic) Expand(req *types.ExpandRequest) (resp *types.ExpandResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
