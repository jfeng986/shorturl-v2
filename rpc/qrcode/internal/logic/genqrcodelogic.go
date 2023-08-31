package logic

import (
	"context"
	"encoding/base64"
	"os"
	"strings"

	"shorturl-v2/rpc/qrcode/internal/svc"
	"shorturl-v2/rpc/qrcode/qrcode"
	"shorturl-v2/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenQrcodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenQrcodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenQrcodeLogic {
	return &GenQrcodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenQrcodeLogic) GenQrcode(in *qrcode.QrcodeRequest) (*qrcode.QrcodeResponse, error) {
	content := in.URL
	err := util.UrlValidation(content)
	if err != nil {
		return nil, err
	}
	filename := content + ".png"
	filename = strings.ReplaceAll(filename, "/", "")
	err = util.GenerateQRCode(content, util.Medium, 256, filename)
	if err != nil {
		return nil, err
	}
	imageData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	base64Str := base64.StdEncoding.EncodeToString(imageData)
	os.Remove(filename)
	resp := &qrcode.QrcodeResponse{
		QrcodeData: base64Str,
	}
	return resp, nil
}
