package logic

import (
	"context"
	"log"
	"net/url"

	"shorturl-v2/api/internal/svc"
	"shorturl-v2/api/internal/types"
	"shorturl-v2/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShortURLLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShortURLLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShortURLLogic {
	return &GetShortURLLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShortURLLogic) GetShortURL(req *types.ShortenRequest) (resp *types.ShortenResponse, err error) {
	originalUrl, customAlias := req.OriginalURL, req.CustomAlias

	err = util.UrlValidation(originalUrl)
	if err != nil {
		return nil, err
	}

	var alias string
	if customAlias == "" {
		alias = util.HashUrl(originalUrl)
	} else {
		alias = customAlias
	}
	shortUrl := "http://127.0.0.1:8888/" + alias
	err = util.UrlValidation(shortUrl)
	if err != nil {
		return nil, err
	}
	parsedShortUrl, err := url.Parse(shortUrl)
	if err != nil {
		return nil, err
	}
	log.Println("parsedUrl:", parsedShortUrl)

	resp = &types.ShortenResponse{
		ShortURL: shortUrl,
	}

	return
}
