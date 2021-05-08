FROM golang:alpine AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=arm

WORKDIR /build
COPY . .
RUN go build -o bin/k8s-status-discord-bot cmd/k8s-status-discord-bot/main.go

FROM alpine:3

RUN apk --no-cache add ca-certificates
WORKDIR /dist
COPY --from=builder /build/bin/k8s-status-discord-bot .
CMD ["/dist/k8s-status-discord-bot"]
