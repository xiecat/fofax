# syntax=docker/dockerfile:1

# ---- build stage -----------------------------------------------------------
FROM golang:1.24-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
COPY vendor ./vendor
COPY cmd ./cmd
COPY internal ./internal
ARG VERSION=dev
RUN CGO_ENABLED=0 go build -mod=vendor -trimpath \
    -ldflags "-s -w -X github.com/xiecat/fofax/internal/cli.FoFaXVersion=${VERSION}" \
    -o /out/fofax ./cmd/fofax

# ---- runtime stage ---------------------------------------------------------
FROM alpine:3.20 AS runtime
RUN adduser -D -u 1000 fofax && mkdir -p /data && chown -R fofax:fofax /data
COPY --from=build /out/fofax /usr/local/bin/fofax
# MCP_TRANSPORT=stdio 是默认值：docker run -i 供本地 MCP 客户端直接拉起。
# 需要端口服务时设 MCP_TRANSPORT=http 并映射 8080。
ENV MCP_TRANSPORT=stdio \
    MCP_HOST=0.0.0.0 \
    MCP_PORT=8080 \
    HOME=/data
USER fofax
WORKDIR /data
VOLUME ["/data"]
EXPOSE 8080
ENTRYPOINT ["fofax"]
