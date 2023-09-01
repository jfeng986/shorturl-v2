package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"shorturl-v2/gateway/internal/config"
	"shorturl-v2/gateway/internal/handler"
	"shorturl-v2/gateway/internal/svc"

	"github.com/joho/godotenv"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/gateway.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// server := rest.MustNewServer(c.RestConf)
	server := rest.MustNewServer(c.RestConf)
	server.Use(corsMiddleware)

	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error fetching current directory:", err)
		} else {
			fmt.Println("Current directory:", cwd)
		}
		if err := godotenv.Load(); err != nil {
			panic("No .env file found")
		}
		allowedOrigin := os.Getenv("ALLOWED_ORIGIN")
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}
