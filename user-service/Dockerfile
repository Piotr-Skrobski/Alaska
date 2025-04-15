# Deploy
FROM golang:1.24-alpine AS builder
ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64
WORKDIR /app
COPY ./user-service/go.mod ./go.mod
COPY ./user-service/go.sum ./go.sum
COPY ./shared-events ./shared-events
RUN go mod download
COPY ./user-service/. ./
RUN go build -o user-service ./cmd/app

# Run
FROM alpine:latest
WORKDIR /
COPY --from=builder /app/user-service .
RUN apk add --no-cache curl
ENTRYPOINT ["/user-service"]