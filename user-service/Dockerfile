# Deploy
FROM golang:1.24-alpine AS builder
ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64
WORKDIR /app

COPY shared-events /app/shared-events
WORKDIR /
COPY user-service/ /app/user-service/
WORKDIR /app/user-service
RUN go build -o user-service ./cmd/app

# Run
FROM alpine:latest
WORKDIR /
COPY --from=builder /app/user-service .
RUN apk add --no-cache curl
ENTRYPOINT ["/user-service"]