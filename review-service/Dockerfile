# Deploy
FROM golang:1.24-alpine AS builder
ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64
WORKDIR /app
COPY ./review-service/go.mod ./go.mod
COPY ./review-service/go.sum ./go.sum
COPY ./shared-events ./shared-events
RUN go mod download
COPY ./review-service/. ./
RUN go build -o review-service ./cmd/app

# Run
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /app/review-service .
USER nonroot:nonroot
ENTRYPOINT ["/review-service"]