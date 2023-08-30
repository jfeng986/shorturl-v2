package logic

import (
	"context"
	"log"
	"net/url"

	"shorturl-v2/rpc/transform/internal/cache"
	"shorturl-v2/rpc/transform/internal/svc"
	"shorturl-v2/rpc/transform/transform"
	"shorturl-v2/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShortenLogic) Shorten(in *transform.ShortenRequest) (*transform.ShortenResponse, error) {
	originalUrl, customAlias := in.OriginalURL, in.CustomAlias

	err := util.UrlValidation(originalUrl)
	if err != nil {
		return nil, err
	}

	var alias string
	if customAlias == "" {
		alias = util.HashUrl(originalUrl)
	} else {
		alias = customAlias
	}
	shortUrl := "http://127.0.0.1:30000/" + alias
	parsedShortUrl, err := url.Parse(shortUrl)
	if err != nil {
		return nil, err
	}
	log.Println("parsedUrl:", parsedShortUrl)
	cache.Put([]byte(shortUrl), []byte(originalUrl))
	resp := &transform.ShortenResponse{
		ShortURL: shortUrl,
	}

	return resp, nil
}
