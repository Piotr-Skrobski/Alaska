# Deploy
FROM golang:1.24-alpine AS builder
ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64
WORKDIR /app
COPY shared-events /app/shared-events
WORKDIR /
COPY movie-service/ /app/movie-service/
WORKDIR /app/movie-service
RUN go build -o movie-service ./cmd/app
# Run
FROM alpine:latest
WORKDIR /
COPY --from=builder /app/movie-service .
RUN apk add --no-cache curl
ENTRYPOINT ["/movie-service"]
