module github.com/Piotr-Skrobski/Alaska/user-service

go 1.24.1

require (
	github.com/Piotr-Skrobski/Alaska/shared-events v0.0.0
	github.com/go-chi/chi/v5 v5.2.2
	github.com/go-redis/redis/v8 v8.11.5
	github.com/lib/pq v1.10.9
	github.com/rabbitmq/amqp091-go v1.10.0
	golang.org/x/crypto v0.37.0
	golang.org/x/net v0.39.0
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
)

replace github.com/Piotr-Skrobski/Alaska/shared-events => ../shared-events
