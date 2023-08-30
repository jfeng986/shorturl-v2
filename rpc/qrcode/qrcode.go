package main

import (
	"flag"
	"fmt"

	"shorturl-v2/rpc/qrcode/internal/config"
	"shorturl-v2/rpc/qrcode/internal/server"
	"shorturl-v2/rpc/qrcode/internal/svc"
	"shorturl-v2/rpc/qrcode/qrcode"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/qrcode.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		qrcode.RegisterQrcoderServer(grpcServer, server.NewQrcoderServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
