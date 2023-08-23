package logic

import (
	"context"

	"shorturl-v2/rpc/transform/internal/svc"
	"shorturl-v2/rpc/transform/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOriginalURLLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOriginalURLLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOriginalURLLogic {
	return &GetOriginalURLLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOriginalURLLogic) GetOriginalURL(in *transform.GetOriginalURLRequest) (*transform.GetOriginalURLResponse, error) {
	// todo: add your logic here and delete this line

	return &transform.GetOriginalURLResponse{}, nil
}
