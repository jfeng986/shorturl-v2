package logic

import (
	"context"
	"errors"
	"log"

	cache "shorturl-v2/rpc/transform/internal/db"
	"shorturl-v2/rpc/transform/internal/svc"
	"shorturl-v2/rpc/transform/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExpandLogic) Expand(in *transform.ExpandRequest) (*transform.ExpandResponse, error) {
	shortURL := in.ShortURL

	originalUrl, ok := cache.Get([]byte(shortURL))
	if ok {
		log.Println(originalUrl)
		resp := &transform.ExpandResponse{
			OriginalURL: originalUrl,
		}
		return resp, nil
	} else {
		return nil, errors.New("short url not found")
	}
}
