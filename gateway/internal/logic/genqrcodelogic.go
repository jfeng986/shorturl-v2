package logic

import (
	"context"

	"shorturl-v2/gateway/internal/svc"
	"shorturl-v2/gateway/internal/types"
	"shorturl-v2/rpc/qrcode/qrcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenQrcodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenQrcodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenQrcodeLogic {
	return &GenQrcodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenQrcodeLogic) GenQrcode(req *types.QrCodeRequest) (*types.QrCodeResponse, error) {
	resp, err := l.svcCtx.Qrcoder.GenQrcode(l.ctx, &qrcode.QrcodeRequest{
		URL: req.URL,
	})
	if err != nil {
		return nil, err
	}
	return &types.QrCodeResponse{
		Qrcode: resp.QrcodeData,
	}, nil
}
