# Deploy
FROM golang:1.24-alpine AS builder
ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64
WORKDIR /app

COPY shared-events /app/shared-events
WORKDIR /
COPY review-service/ /app/review-service/
WORKDIR /app/review-service
RUN go build -o review-service ./cmd/app

# Run
FROM alpine:latest
WORKDIR /
COPY --from=builder /app/review-service .
RUN apk add --no-cache curl
ENTRYPOINT ["/review-service"]