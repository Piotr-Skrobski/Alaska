# CLAUDE.md - Alaska Movie Review Platform

This document provides a comprehensive guide to the Alaska codebase for AI assistants working on this project.

## Table of Contents
1. [Architecture Overview](#architecture-overview)
2. [Service Details](#service-details)
3. [Development Environment](#development-environment)
4. [Code Structure & Conventions](#code-structure--conventions)
5. [Database Schema](#database-schema)
6. [API Patterns](#api-patterns)
7. [Development Workflows](#development-workflows)
8. [Testing Guidelines](#testing-guidelines)
9. [Common Tasks](#common-tasks)
10. [Important Notes](#important-notes)

---

## Architecture Overview

**Type**: Microservices architecture in a monorepo structure
**Purpose**: Movie review platform with user authentication and LLM integration

### Services Map
```
┌──────────────────────────────────────────────────────────────┐
│                         Nginx (Port 80)                       │
│           Reverse proxy for Frontend & Ollama                │
└────────────┬──────────────────────────────────┬──────────────┘
             │                                  │
    ┌────────▼─────────┐              ┌────────▼─────────┐
    │    Frontend      │              │     Ollama       │
    │  (Vue 3 + Vite)  │              │   (LLM Service)  │
    │    Port 8081     │              │   Port 11434     │
    └────────┬─────────┘              └──────────────────┘
             │
    ┌────────▼─────────┐
    │     Gateway      │
    │  (API Gateway)   │
    │    Port 8080     │
    └────────┬─────────┘
             │
    ┌────────┴────────────────────────────────┐
    │                                         │
┌───▼────────────┐  ┌──────────────┐  ┌──────▼────────┐
│  User Service  │  │ Movie Service│  │Review Service │
│   Port 10001   │  │  Port 10002  │  │  Port 10003   │
└───┬────────┬───┘  └──┬───────┬───┘  └──┬────────────┘
    │        │         │       │         │
    │   ┌────▼─────────▼───────▼─────────▼────┐
    │   │        RabbitMQ (Event Bus)         │
    │   │  Port 5672 (AMQP), 15672 (Mgmt)    │
    │   └─────────────────────────────────────┘
    │
┌───▼──────────┐  ┌──────────────┐  ┌──────────────┐
│  PostgreSQL  │  │   MongoDB    │  │    Redis     │
│  Port 5432   │  │  Port 27017  │  │  Port 6379   │
│ (Users, Rvws)│  │   (Movies)   │  │  (Sessions)  │
└──────────────┘  └──────────────┘  └──────────────┘
```

### Network Architecture
- **alaska_network**: Main network connecting all application services
- **message_broker_network**: Internal network for RabbitMQ (isolated)
- **mongodb_network**: Internal network for MongoDB and Mongo Express
- **postgres_network**: Internal network for PostgreSQL and pgAdmin
- **redis_network**: Internal network for Redis (isolated)

---

## Service Details

### 1. Gateway Service (`gateway-service/`)
**Tech Stack**: Go 1.24.1, Chi router v5.2.1
**Port**: 8080
**Purpose**: Reverse proxy / API Gateway

**Key Responsibilities**:
- Routes requests to appropriate backend services
- CORS handling for frontend origins
- Health check endpoint at `/health`

**Routing Rules**:
- `/users/*` → user-service
- `/movies/*` → movie-service
- `/reviews/*` → review-service

**Important Files**:
- `cmd/app/main.go`: Entry point
- `internal/handlers/`: HTTP proxy handlers
- `internal/utils/Config.go`: Configuration

### 2. User Service (`user-service/`)
**Tech Stack**: Go 1.24.1, Chi router, PostgreSQL, Redis, RabbitMQ
**Port**: 10001 (typical)
**Database**: PostgreSQL (table: `users`)
**Cache**: Redis (sessions with 6-hour TTL)

**Key Responsibilities**:
- User registration and authentication
- Session management (HTTP-only cookies)
- Password hashing with bcrypt
- Discord OAuth integration
- User deletion with event publishing

**Environment Variables**:
- `POSTGRES_URI`: PostgreSQL connection string
- `REDIS_URI`: Redis connection string (format: `redis://:<password>@<host>:6379`)
- `RABBITMQ_URI`: RabbitMQ connection string
- `JWT_SECRET`: Secret for JWT signing
- `DISCORD_CLIENT_ID`, `DISCORD_CLIENT_SECRET`: OAuth credentials
- `PORT`: Service port (default: 10002)

**API Endpoints**:
- `POST /users/register`: User registration
- `POST /users/login`: User login (creates session)
- `GET /users/me`: Get current user (requires auth)
- `POST /users/logout`: Logout (destroys session)
- `POST /users/delete`: Delete account (requires auth, publishes event)
- `GET /users/health`: Health check

**Events Published**:
- `UserRegisteredEvent`: On successful registration
- `UserDeleted`: On account deletion

### 3. Movie Service (`movie-service/`)
**Tech Stack**: Go 1.24.1, Chi router, MongoDB, RabbitMQ
**Port**: 10002 (typical)
**Database**: MongoDB (collection: `movies`)
**External API**: OMDb API for movie data

**Key Responsibilities**:
- Fetch movie data from OMDb API
- Cache movies in MongoDB
- Movie CRUD operations
- Listen for user deletion events to cascade deletions

**Environment Variables**:
- `MONGO_URI`: MongoDB connection string
- `RABBITMQ_URI`: RabbitMQ connection string
- `OMDB_API_KEY`: OMDb API key
- `PORT`: Service port (default: 10002)

**Data Flow**:
1. Request for movie comes in
2. Check MongoDB cache
3. If not found, fetch from OMDb API
4. Store in MongoDB
5. Return to client

**API Endpoints**:
- `GET /movies/:id`: Get movie by ID (IMDb ID)
- `POST /movies`: Create/update movie
- `DELETE /movies/:id`: Delete movie
- `GET /movies/health`: Health check

**Events Consumed**:
- `UserDeleted`: Triggers cascade deletion of related data

**Events Published**:
- `MovieDeletedEvent`: When a movie is deleted

### 4. Review Service (`review-service/`)
**Tech Stack**: Go 1.24.1, Chi router, PostgreSQL
**Port**: 10003 (typical)
**Database**: PostgreSQL (table: `reviews`)

**Key Responsibilities**:
- Create and retrieve movie reviews
- Associate reviews with users and movies
- Cascade delete reviews when users or movies are deleted

**Environment Variables**:
- `POSTGRES_URI`: PostgreSQL connection string
- `RABBITMQ_URI`: RabbitMQ connection string
- `PORT`: Service port (default: 10002)

**Database Schema**:
```sql
CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    movie_id VARCHAR(255) NOT NULL,
    rating INTEGER CHECK (rating >= 1 AND rating <= 10),
    comment TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

**API Endpoints**:
- `POST /reviews`: Create a review
- `GET /reviews/movie/:movieId`: Get reviews for a movie
- `GET /reviews/user/:userId`: Get reviews by a user
- `GET /reviews/health`: Health check

**Events Consumed**:
- `UserDeleted`: Cascades review deletion
- `MovieDeletedEvent`: Cascades review deletion

### 5. Frontend (`frontend/`)
**Tech Stack**: Vue 3.5.13, Vite 6.2.4, Vue Router 4.5.0, Axios 1.8.4
**Dev Port**: 5173 (Vite dev server)
**Prod Port**: 8081 (served via Nginx)

**Key Features**:
- Movie search and browsing
- User registration and authentication
- Review creation and viewing
- Profile page
- LLM integration (recent addition)

**Build Commands**:
- `npm run dev`: Start Vite dev server
- `npm run build`: Build for production
- `npm run preview`: Preview production build

**Directory Structure**:
```
frontend/
├── src/
│   ├── components/      # Reusable Vue components
│   ├── views/          # Page-level components
│   ├── router/         # Vue Router configuration
│   ├── services/       # API service layer (Axios)
│   ├── App.vue
│   └── main.js
├── public/
├── package.json
├── vite.config.js
└── Dockerfile
```

**API Communication**:
- Uses Axios for HTTP requests
- Base URL configured to point to Gateway service
- Session cookies automatically sent with requests

### 6. Shared Events (`shared-events/`)
**Tech Stack**: Go 1.24.1
**Type**: Shared library (Go module)

**Purpose**: Type-safe event definitions for inter-service communication

**Event Definitions**:
```go
// UserRegisteredEvent - Published by user-service on registration
type UserRegisteredEvent struct {
    UserID   int    `json:"user_id"`
    Username string `json:"username"`
    Email    string `json:"email"`
}

// UserDeleted - Published by user-service on account deletion
type UserDeleted struct {
    UserID int `json:"user_id"`
}

// ReviewCreatedEvent - Published by review-service
type ReviewCreatedEvent struct {
    ReviewID int    `json:"review_id"`
    UserID   int    `json:"user_id"`
    MovieID  string `json:"movie_id"`
    Rating   int    `json:"rating"`
}

// MovieDeletedEvent - Published by movie-service
type MovieDeletedEvent struct {
    MovieID string `json:"movie_id"`
}
```

**Usage**: Services import this module and use these types for RabbitMQ message payloads

---

## Development Environment

### Prerequisites
- Docker and Docker Compose
- Go 1.24+ (for local development)
- Node.js (current version) (for frontend development)
- Git

### Environment Variables

Create a `.env` file in the root directory with the following variables:

```bash
# Database Credentials
POSTGRES_USER=your_postgres_user
POSTGRES_PASSWORD=your_postgres_password
POSTGRES_DB=alaska_db
POSTGRES_URL=postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@postgres:5432/${POSTGRES_DB}?sslmode=require

# MongoDB
MONGO_USER=your_mongo_user
MONGO_PASSWORD=your_mongo_password
MONGO_DATABASE=movies_db
MONGO_URI=mongodb://${MONGO_USER}:${MONGO_PASSWORD}@mongodb:27017/${MONGO_DATABASE}?authSource=admin

# Redis
REDIS_PASSWORD=your_redis_password
REDIS_URI=redis://:${REDIS_PASSWORD}@redis:6379

# RabbitMQ
RABBITMQ_USER=your_rabbitmq_user
RABBITMQ_PASSWORD=your_rabbitmq_password
RABBITMQ_URI=amqp://${RABBITMQ_USER}:${RABBITMQ_PASSWORD}@rabbitmq:5672/

# Service Ports
GATEWAY_PORT=8080
USER_SERVICE_PORT=10001
MOVIE_SERVICE_PORT=10002
REVIEW_SERVICE_PORT=10003
PORT=10002

# Application Secrets
JWT_SECRET=your_jwt_secret_key
OMDB_API_KEY=your_omdb_api_key

# Discord OAuth (optional)
DISCORD_CLIENT_ID=your_discord_client_id
DISCORD_CLIENT_SECRET=your_discord_client_secret

# Admin Panels
MONGO_EXPRESS_USER=admin
MONGO_EXPRESS_PASSWORD=admin_password
PGADMIN_EMAIL=admin@admin.com
PGADMIN_PASSWORD=admin_password
```

**IMPORTANT**: Never commit `.env` file to git (already in `.gitignore`)

### Starting the Development Environment

**Full Stack with Docker Compose**:
```bash
docker-compose up --build
```

This starts:
- All 4 backend services
- Frontend (production build)
- PostgreSQL + pgAdmin (http://localhost:5433)
- MongoDB + Mongo Express (http://localhost:27018)
- Redis
- RabbitMQ + Management UI (http://localhost:15672)
- Ollama LLM service
- Nginx reverse proxy (http://localhost:80)

**Frontend Development Mode** (hot-reload):
```bash
cd frontend
npm install
npm run dev
# Frontend available at http://localhost:5173
```

**Individual Service Development** (requires infrastructure running):
```bash
# Start infrastructure only
docker-compose up postgres mongodb redis rabbitmq

# Run a service locally
cd user-service
go run cmd/app/main.go
```

### Health Checks

All services expose health check endpoints:
- Gateway: `http://localhost:8080/health`
- User Service: `http://localhost:10001/users/health`
- Movie Service: `http://localhost:10002/movies/health`
- Review Service: `http://localhost:10003/reviews/health`

---

## Code Structure & Conventions

### Go Services - Standard Layout

All Go services follow this structure:

```
<service-name>/
├── cmd/
│   └── app/
│       └── main.go              # Entry point, wires dependencies
├── internal/                     # Private application code
│   ├── controllers/              # HTTP request handlers
│   │   └── *Controller.go        # PascalCase, exported
│   ├── services/                 # Business logic layer
│   │   └── *Service.go           # PascalCase, exported
│   ├── repositories/             # Data access layer
│   │   └── *Repository.go        # PascalCase, exported
│   ├── models/                   # Domain models
│   │   └── *.go                  # PascalCase, exported types
│   ├── dtos/                     # Data transfer objects
│   │   └── *.go                  # Request/Response structs
│   ├── routers/                  # Route configuration
│   │   └── Router.go             # Chi router setup
│   └── utils/                    # Utilities and config
│       └── Config.go             # Configuration loading
├── Dockerfile                    # Multi-stage build
├── go.mod                        # Dependencies
└── go.sum
```

### Naming Conventions

**Go Code**:
- Files: PascalCase (e.g., `UserController.go`, `UserService.go`)
- Types: PascalCase, exported (e.g., `UserController`, `LoginRequest`)
- Interfaces: PascalCase with 'er' suffix when appropriate (e.g., `UserRepository`)
- Functions: PascalCase for exported, camelCase for private
- Variables: camelCase
- Constants: PascalCase or ALL_CAPS depending on context

**Vue/JavaScript**:
- Components: PascalCase (e.g., `MovieCard.vue`, `ProfilePage.vue`)
- Files in views/: PascalCase (e.g., `HomePage.vue`)
- Services: camelCase (e.g., `userService.js`)
- Variables/functions: camelCase

### Dependency Injection Pattern

All Go services use constructor injection:

```go
// Example from user-service/cmd/app/main.go
func main() {
    config := utils.LoadConfig()

    // Infrastructure
    db := setupDatabase(config)
    redisClient := setupRedis(config)
    rabbitConn := setupRabbitMQ(config)

    // Repositories
    userRepo := repositories.NewUserRepository(db)

    // Services
    userService := services.NewUserService(userRepo, redisClient, rabbitConn)

    // Controllers
    userController := controllers.NewUserController(userService)

    // Routes
    router := routers.SetupRouter(userController)

    http.ListenAndServe(":"+config.Port, router)
}
```

### Layered Architecture

**Request Flow**:
```
HTTP Request
    ↓
Controller (validation, HTTP concerns)
    ↓
Service (business logic, orchestration)
    ↓
Repository (database operations)
    ↓
Database
```

**Key Principles**:
- Controllers should NOT contain business logic
- Services should NOT know about HTTP (no `http.Request` or `http.ResponseWriter`)
- Repositories should only handle data access
- Use DTOs for request/response, Models for domain entities

### Error Handling

**Go Services**:
```go
// Controllers return errors, don't panic
func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
    var req dtos.RegisterRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "invalid request", http.StatusBadRequest)
        return
    }

    if err := uc.UserService.Register(r.Context(), req); err != nil {
        http.Error(w, "failed to create user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}
```

**Custom Errors** (movie-service example):
```go
// internal/errors/errors.go
type NotFoundError struct {
    Message string
}

func (e *NotFoundError) Error() string {
    return e.Message
}
```

### Logging Conventions

Services use emoji indicators for log messages:
- ✅ Success/startup messages
- 🚀 Service start
- ❌ Errors (though not consistently used)

Example:
```go
log.Println("🚀 User Service starting on port", config.Port)
log.Println("✅ Connected to PostgreSQL")
```

**Note**: This is informal logging. Consider structured logging for production.

---

## Database Schema

### PostgreSQL (User Service)

**Table: `users`**
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL  -- Soft delete support
);
```

### PostgreSQL (Review Service)

**Table: `reviews`**
```sql
CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    movie_id VARCHAR(255) NOT NULL,  -- IMDb ID
    rating INTEGER CHECK (rating >= 1 AND rating <= 10),
    comment TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

**Note**: No foreign key constraints (microservices pattern)

### MongoDB (Movie Service)

**Collection: `movies`**

Document structure (from OMDb API):
```json
{
    "_id": "tt1234567",  // IMDb ID
    "Title": "Movie Title",
    "Year": "2020",
    "Rated": "PG-13",
    "Released": "01 Jan 2020",
    "Runtime": "120 min",
    "Genre": "Action, Drama",
    "Director": "Director Name",
    "Writer": "Writer Name",
    "Actors": "Actor 1, Actor 2",
    "Plot": "Plot summary...",
    "Language": "English",
    "Country": "USA",
    "Awards": "Won 2 Oscars",
    "Poster": "https://...",
    "Ratings": [...],
    "Metascore": "85",
    "imdbRating": "8.5",
    "imdbVotes": "1,000,000",
    "imdbID": "tt1234567",
    "Type": "movie",
    "DVD": "01 May 2020",
    "BoxOffice": "$100,000,000",
    "Production": "Production Company",
    "Website": "https://...",
    "Response": "True"
}
```

### Redis (User Service)

**Key Pattern**: `session:<session_id>`
**Value**: User ID (integer as string)
**TTL**: 6 hours (21600 seconds)

---

## API Patterns

### Authentication Flow

1. **Registration**:
```
POST /users/register
Body: { "username": "...", "email": "...", "password": "..." }
Response: 201 Created
Event: UserRegisteredEvent published to RabbitMQ
```

2. **Login**:
```
POST /users/login
Body: { "username": "...", "password": "..." }
Response: 200 OK + Set-Cookie: session_id=...; HttpOnly
Redis: session:<id> → user_id (6h TTL)
```

3. **Authenticated Request**:
```
GET /users/me
Cookie: session_id=...
Process:
  - Middleware checks cookie
  - Looks up session in Redis
  - Injects user_id into context
  - Controller retrieves user
```

4. **Logout**:
```
POST /users/logout
Cookie: session_id=...
Process:
  - Delete session from Redis
  - Clear cookie
```

### CORS Configuration

Gateway service handles CORS for these origins:
- `http://localhost:5173` (Vite dev server)
- `http://localhost:8081` (Production frontend)
- `http://localhost:80` (Nginx)

Allowed methods: GET, POST, PUT, DELETE, OPTIONS
Allowed headers: Content-Type, Authorization
Credentials: Allowed

### Event-Driven Communication

**Publishing Events**:
```go
// In user-service
event := events.UserDeleted{UserID: userID}
payload, _ := json.Marshal(event)
channel.Publish("", "user.deleted", false, false, amqp.Publishing{
    ContentType: "application/json",
    Body:        payload,
})
```

**Consuming Events**:
```go
// In review-service
msgs, _ := channel.Consume("user.deleted", "", true, false, false, false, nil)
for msg := range msgs {
    var event events.UserDeleted
    json.Unmarshal(msg.Body, &event)
    // Handle event (e.g., delete user's reviews)
}
```

**Queue Names**:
- `user.registered`
- `user.deleted`
- `review.created`
- `movie.deleted`

---

## Development Workflows

### Adding a New Endpoint

1. **Define DTO** (if needed):
```go
// internal/dtos/request.go
type CreateReviewRequest struct {
    MovieID string `json:"movie_id"`
    Rating  int    `json:"rating"`
    Comment string `json:"comment"`
}
```

2. **Add Repository Method**:
```go
// internal/repositories/ReviewRepository.go
func (r *ReviewRepository) Create(ctx context.Context, review models.Review) error {
    query := `INSERT INTO reviews (user_id, movie_id, rating, comment) VALUES ($1, $2, $3, $4)`
    _, err := r.db.ExecContext(ctx, query, review.UserID, review.MovieID, review.Rating, review.Comment)
    return err
}
```

3. **Add Service Method**:
```go
// internal/services/ReviewService.go
func (s *ReviewService) CreateReview(ctx context.Context, req dtos.CreateReviewRequest, userID int) error {
    review := models.Review{
        UserID:  userID,
        MovieID: req.MovieID,
        Rating:  req.Rating,
        Comment: req.Comment,
    }
    return s.repo.Create(ctx, review)
}
```

4. **Add Controller Handler**:
```go
// internal/controllers/ReviewController.go
func (c *ReviewController) Create(w http.ResponseWriter, r *http.Request) {
    var req dtos.CreateReviewRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "invalid request", http.StatusBadRequest)
        return
    }

    userID := r.Context().Value("user_id").(int)

    if err := c.service.CreateReview(r.Context(), req, userID); err != nil {
        http.Error(w, "failed to create review", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}
```

5. **Register Route**:
```go
// internal/routers/Router.go or Controller
r.Post("/reviews", controller.Create)
```

### Adding a New Event

1. **Define Event in shared-events**:
```go
// shared-events/events/new_event.go
package events

type NewEvent struct {
    Field1 string `json:"field1"`
    Field2 int    `json:"field2"`
}
```

2. **Update go.mod in consuming services**:
```bash
cd <service>
go mod tidy
```

3. **Publish Event** (in producing service):
```go
event := events.NewEvent{Field1: "value", Field2: 123}
payload, _ := json.Marshal(event)
channel.Publish("", "event.name", false, false, amqp.Publishing{
    ContentType: "application/json",
    Body:        payload,
})
```

4. **Consume Event** (in consuming service):
```go
// In service initialization
msgs, _ := channel.Consume("event.name", "", true, false, false, false, nil)
go func() {
    for msg := range msgs {
        var event events.NewEvent
        json.Unmarshal(msg.Body, &event)
        // Handle event
    }
}()
```

### Database Migrations

**Current Approach**: Code-first auto-migration on service startup

Services check and create tables if they don't exist in `main.go`:
```go
// Example from user-service
func createUsersTable(db *sql.DB) {
    query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(255) UNIQUE NOT NULL,
        email VARCHAR(255) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP NULL
    )`
    db.Exec(query)
}
```

**For Schema Changes**:
1. Modify the CREATE TABLE query in `cmd/app/main.go`
2. For existing deployments, manually run ALTER TABLE or use migration tools
3. **TODO**: Consider adopting a proper migration tool (e.g., golang-migrate)

### Git Commit Conventions

Based on git history, the project uses conventional commits:

**Format**: `<type>(<scope>): <description>`

**Types**:
- `feat`: New feature
- `fix`: Bug fix
- `refactor`: Code refactoring
- `build`: Build system or dependency changes
- `chore`: Maintenance tasks

**Scopes**:
- `Frontend`, `Docker`, `Infra`, `User`, `Movie`, `Review`, `Gateway`, `Event`

**Examples**:
```
feat(Frontend): Add LLM integration
fix(Infra): Fix CORS issues
refactor(User): Refactor User Controller into User Service
feat(Docker): Add Ollama local model to Docker
```

---

## Testing Guidelines

**Current State**: No automated tests exist in the codebase.

**Recommended Approach for Future Tests**:

### Unit Tests (Go)
```go
// user-service/internal/services/UserService_test.go
package services_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

func TestUserService_Register(t *testing.T) {
    // Arrange
    mockRepo := new(MockUserRepository)
    service := NewUserService(mockRepo, nil, nil)

    // Act
    err := service.Register(context.Background(), dtos.RegisterRequest{
        Username: "testuser",
        Email: "test@example.com",
        Password: "password123",
    })

    // Assert
    assert.NoError(t, err)
    mockRepo.AssertExpectations(t)
}
```

### Integration Tests
Test database interactions with test containers or test databases.

### E2E Tests (Frontend)
Consider Playwright or Cypress for frontend testing.

### API Tests
Use `net/http/httptest` for testing HTTP handlers:
```go
func TestUserController_Register(t *testing.T) {
    req := httptest.NewRequest("POST", "/users/register", body)
    w := httptest.NewRecorder()

    controller.Register(w, req)

    assert.Equal(t, http.StatusCreated, w.Code)
}
```

---

## Common Tasks

### Adding a New Microservice

1. Create service directory: `<service-name>/`
2. Initialize Go module: `go mod init github.com/Piotr-Skrobski/Alaska/<service-name>`
3. Create standard directory structure (cmd/, internal/, Dockerfile)
4. Add shared-events dependency if needed
5. Create service configuration in docker-compose.yml
6. Add routing in gateway-service
7. Update this document

### Debugging

**Check Service Logs**:
```bash
docker-compose logs -f <service-name>
docker-compose logs -f user-service
```

**Check Database**:
- PostgreSQL: http://localhost:5433 (pgAdmin)
- MongoDB: http://localhost:27018 (Mongo Express)
- Redis: Use redis-cli or RedisInsight

**Check Message Queue**:
- RabbitMQ Management: http://localhost:15672
- Default credentials: Set in .env

**Test Health Endpoints**:
```bash
curl http://localhost:8080/health
curl http://localhost:10001/users/health
curl http://localhost:10002/movies/health
curl http://localhost:10003/reviews/health
```

### Environment Variable Updates

1. Add to `.env` file (local)
2. Add to `docker-compose.yml` in service's `environment:` section
3. Update `internal/utils/Config.go` in the service
4. Restart service: `docker-compose up -d --build <service-name>`

### Updating Dependencies

**Go Services**:
```bash
cd <service-name>
go get -u <package>
go mod tidy
```

**Frontend**:
```bash
cd frontend
npm update <package>
# or
npm install <package>@latest
```

**Rebuild containers after dependency changes**:
```bash
docker-compose up -d --build <service-name>
```

---

## Important Notes

### Security Considerations

1. **Passwords**: Hashed with bcrypt (default cost factor 10)
2. **Sessions**:
   - Stored in Redis with 6-hour TTL
   - HTTP-only cookies (not accessible via JavaScript)
   - No CSRF protection implemented yet
3. **CORS**: Configured but review allowed origins regularly
4. **SQL Injection**: Using parameterized queries ($1, $2 placeholders)
5. **Secrets Management**: Environment variables (consider secrets manager for production)
6. **No Rate Limiting**: Consider adding rate limiting middleware
7. **No Input Validation Library**: Manual validation (consider validator package)

### Performance Considerations

1. **Resource Limits**: Docker Compose sets CPU (0.5) and memory (512MB) limits per service
2. **Database Connections**: No connection pooling configuration visible
3. **Caching**: Movie data cached in MongoDB, user sessions in Redis
4. **No CDN**: Static assets served directly from Nginx
5. **MongoDB Cache Size**: Limited to 1GB WiredTiger cache

### Known Issues and TODOs

Based on code review:
1. No automated tests
2. No API documentation (consider Swagger/OpenAPI)
3. Code comments mention "Refactor later" in several places
4. No proper migration system for databases
5. Logging is informal (emoji-based, not structured)
6. No observability (metrics, tracing)
7. Error messages could be more specific
8. No input validation framework
9. Some environment fallbacks but not comprehensive

### Scaling Considerations

Current architecture supports horizontal scaling with considerations:
- **Stateless Services**: User service uses Redis for sessions (shared state)
- **Database**: Single PostgreSQL and MongoDB instances (bottleneck)
- **Message Queue**: Single RabbitMQ instance
- **Session Affinity**: Not required (Redis-backed sessions)

For production scaling:
1. Set up database replicas (read replicas)
2. Use managed RabbitMQ cluster
3. Consider Redis cluster for sessions
4. Add load balancer in front of gateway
5. Implement service discovery (Consul, etcd)

### External Dependencies

1. **OMDb API**: Requires API key, rate limits apply
2. **Discord OAuth**: Optional, requires application setup
3. **Ollama**: Self-hosted LLM, pulls llama3 model on startup

### File Paths Reference

Quick reference for common files:

```
/home/user/Alaska/
├── docker-compose.yml              # Orchestration config
├── .env                            # Environment variables (gitignored)
├── .gitignore                      # Git ignore rules
├── ollama-entrypoint.sh           # Ollama startup script
│
├── gateway-service/
│   ├── cmd/app/main.go
│   └── internal/handlers/
│
├── user-service/
│   ├── cmd/app/main.go
│   └── internal/
│       ├── controllers/UserController.go
│       ├── services/UserService.go
│       ├── repositories/UserRepository.go
│       ├── models/User.go
│       └── dtos/
│
├── movie-service/
│   ├── cmd/app/main.go
│   └── internal/
│       ├── controllers/MovieController.go
│       ├── services/MovieService.go
│       └── repositories/MovieRepository.go
│
├── review-service/
│   ├── cmd/app/main.go
│   └── internal/
│       ├── controllers/ReviewController.go
│       ├── services/ReviewService.go
│       └── repositories/ReviewRepository.go
│
├── frontend/
│   ├── src/
│   │   ├── components/
│   │   ├── views/
│   │   ├── router/
│   │   ├── services/
│   │   └── App.vue
│   ├── package.json
│   └── vite.config.js
│
└── shared-events/
    ├── events/
    │   ├── user_events.go
    │   ├── review_events.go
    │   └── movie_events.go
    └── go.mod
```

---

## Working with AI Assistants

### Best Practices

1. **Always check existing patterns** before implementing new features
2. **Follow the layered architecture** (Controller → Service → Repository)
3. **Use dependency injection** consistently
4. **Maintain event-driven communication** for cross-service operations
5. **Update this document** when adding significant new patterns or services
6. **Test locally** before committing (at minimum, verify service starts)

### Quick Checklist for Changes

- [ ] Does it follow the existing code structure?
- [ ] Are errors handled properly?
- [ ] Are environment variables documented?
- [ ] Is the change reflected in docker-compose.yml if needed?
- [ ] Will this require database migrations?
- [ ] Does this affect other services (events)?
- [ ] Is CLAUDE.md updated with new patterns?
- [ ] Git commit follows conventional commit format?

### When to Ask for Clarification

- Unclear business logic requirements
- Authentication/authorization requirements
- Performance requirements
- Deployment environment constraints
- External API integration specifics
- Data retention policies

---

## Useful Commands

```bash
# Start everything
docker-compose up --build

# Start specific services
docker-compose up postgres mongodb redis rabbitmq

# Rebuild single service
docker-compose up -d --build user-service

# View logs
docker-compose logs -f user-service

# Stop everything
docker-compose down

# Stop and remove volumes (DESTRUCTIVE)
docker-compose down -v

# Check running services
docker-compose ps

# Execute command in container
docker-compose exec postgres psql -U your_user -d alaska_db

# Frontend development
cd frontend && npm run dev

# Run Go service locally (requires infrastructure)
cd user-service && go run cmd/app/main.go

# Update Go dependencies
cd user-service && go mod tidy

# Format Go code
go fmt ./...

# Build Go service locally
go build -o bin/service ./cmd/app
```

---

**Document Version**: 1.0
**Last Updated**: 2025-11-15
**Maintained By**: AI Assistant (Claude)

For questions or updates to this document, please open an issue or submit a pull request.
