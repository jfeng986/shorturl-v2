package logic

import (
	"context"

	"shorturl-v2/api/internal/svc"
	"shorturl-v2/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOriginalURLLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOriginalURLLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOriginalURLLogic {
	return &GetOriginalURLLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOriginalURLLogic) GetOriginalURL(req *types.GetOriginalURLRequest) (resp *types.GetOriginalURLResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
