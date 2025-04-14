# Deploy
FROM golang:1.24-alpine AS builder
ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o gateway-service ./cmd/app

# Run
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /app/gateway-service .
USER nonroot:nonroot
ENTRYPOINT ["/gateway-service"]