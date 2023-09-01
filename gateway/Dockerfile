FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build
COPY . .
RUN sh -c "[ -f go.mod ]" || exit
COPY gateway/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/gateway gateway/gateway.go

FROM alpine

RUN echo -e http://mirrors.ustc.edu.cn/alpine/v3.13/main/ > /etc/apk/repositories
RUN cat /etc/apk/repositories
RUN apk update --no-cache
RUN apk add --no-cache ca-certificates

EXPOSE 30000

WORKDIR /app
COPY --from=builder /app/gateway /app/gateway
COPY --from=builder /app/etc /app/etc

CMD ["./gateway", "-f", "etc/gateway.yaml"]
