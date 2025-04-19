module github.com/Piotr-Skrobski/Alaska/review-service

go 1.24.1

require (
	github.com/go-chi/chi/v5 v5.2.1
	github.com/Piotr-Skrobski/Alaska/shared-events v0.0.0
)

require github.com/lib/pq v1.10.9 // indirect

replace github.com/Piotr-Skrobski/Alaska/shared-events => ../shared-events
