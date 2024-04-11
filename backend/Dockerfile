FROM golang:1.22.1-alpine AS base

WORKDIR /app

FROM base AS dev

RUN go install golang.org/x/tools/cmd/goimports@v0.19 && \
    go install github.com/vorlif/xspreak@v0.10.1
