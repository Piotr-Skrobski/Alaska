# Alaska Movie Review Platform

A modern, microservices-based movie review platform with user authentication, movie data integration, and AI-powered features.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Architecture](#architecture)
- [Tech Stack](#tech-stack)
- [Prerequisites](#prerequisites)
- [Quick Start](#quick-start)
- [Environment Configuration](#environment-configuration)
- [Development](#development)
- [Project Structure](#project-structure)
- [API Documentation](#api-documentation)
- [Services](#services)
- [Contributing](#contributing)
- [License](#license)

## Overview

Alaska is a full-stack movie review platform built with a microservices architecture. Users can search for movies, read and write reviews, and interact with an integrated LLM (Large Language Model) for enhanced movie recommendations and insights.

## Features

- **User Management**
  - User registration and authentication
  - Session-based authentication with HTTP-only cookies
  - Discord OAuth integration
  - Secure password hashing with bcrypt

- **Movie Database**
  - Integration with OMDb API for comprehensive movie data
  - MongoDB caching for fast retrieval
  - Detailed movie information including ratings, cast, plot, and more

- **Review System**
  - Create and read movie reviews
  - Rating system (1-10 scale)
  - User-specific and movie-specific review queries

- **AI Integration**
  - Local LLM integration using Ollama (llama3 model)
  - AI-powered movie insights and recommendations

- **Modern Architecture**
  - Microservices-based design
  - Event-driven communication with RabbitMQ
  - RESTful API Gateway
  - Responsive Vue.js frontend

## Architecture

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

### Architecture Highlights

- **API Gateway Pattern**: Central entry point for all client requests
- **Event-Driven Communication**: Services communicate asynchronously via RabbitMQ
- **Database per Service**: Each service owns its data (PostgreSQL, MongoDB)
- **Caching Layer**: Redis for session management, MongoDB for movie data caching
- **Reverse Proxy**: Nginx for efficient static file serving and routing

## Tech Stack

### Backend
- **Language**: Go 1.24.1
- **Web Framework**: Chi Router v5.2.1
- **Databases**: PostgreSQL 17, MongoDB 8.0
- **Cache**: Redis 7.4
- **Message Broker**: RabbitMQ 4.0
- **Authentication**: bcrypt, session-based with Redis

### Frontend
- **Framework**: Vue 3.5.13
- **Build Tool**: Vite 6.2.4
- **Router**: Vue Router 4.5.0
- **HTTP Client**: Axios 1.8.4

### AI/ML
- **LLM Platform**: Ollama
- **Model**: llama3

### Infrastructure
- **Containerization**: Docker & Docker Compose
- **Reverse Proxy**: Nginx 1.27
- **Admin Tools**: pgAdmin, Mongo Express

## Prerequisites

Before you begin, ensure you have the following installed:

- **Docker** (version 20.10 or higher)
- **Docker Compose** (version 2.0 or higher)
- **Git**

Optional (for local development without Docker):
- **Go** 1.24+
- **Node.js** (current LTS version)
- **PostgreSQL** 17
- **MongoDB** 8.0
- **Redis** 7.4
- **RabbitMQ** 4.0

## Quick Start

### 1. Clone the Repository

```bash
git clone https://github.com/Piotr-Skrobski/Alaska.git
cd Alaska
```

### 2. Set Up Environment Variables

Create a `.env` file in the root directory:

```bash
# Database Credentials
POSTGRES_USER=alaska_user
POSTGRES_PASSWORD=your_secure_password
POSTGRES_DB=alaska_db
POSTGRES_URL=postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@postgres:5432/${POSTGRES_DB}?sslmode=require

# MongoDB
MONGO_USER=mongo_user
MONGO_PASSWORD=your_mongo_password
MONGO_DATABASE=movies_db
MONGO_URI=mongodb://${MONGO_USER}:${MONGO_PASSWORD}@mongodb:27017/${MONGO_DATABASE}?authSource=admin

# Redis
REDIS_PASSWORD=your_redis_password
REDIS_URI=redis://:${REDIS_PASSWORD}@redis:6379

# RabbitMQ
RABBITMQ_USER=rabbitmq_user
RABBITMQ_PASSWORD=your_rabbitmq_password
RABBITMQ_URI=amqp://${RABBITMQ_USER}:${RABBITMQ_PASSWORD}@rabbitmq:5672/

# Service Ports
GATEWAY_PORT=8080
USER_SERVICE_PORT=10001
MOVIE_SERVICE_PORT=10002
REVIEW_SERVICE_PORT=10003
PORT=10002

# Application Secrets
JWT_SECRET=your_jwt_secret_key_min_32_chars
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

**Note**:
- Get your free OMDb API key from [http://www.omdbapi.com/apikey.aspx](http://www.omdbapi.com/apikey.aspx)
- Replace all passwords with secure values
- Never commit the `.env` file to version control

### 3. Start the Application

```bash
docker-compose up --build
```

This command will:
- Build all microservices
- Start all databases and infrastructure
- Launch the frontend
- Initialize the Ollama LLM service

### 4. Access the Application

- **Frontend**: [http://localhost](http://localhost) or [http://localhost:80](http://localhost:80)
- **API Gateway**: [http://localhost:8080](http://localhost:8080)
- **pgAdmin**: [http://localhost:5433](http://localhost:5433)
- **Mongo Express**: [http://localhost:27018](http://localhost:27018)
- **RabbitMQ Management**: [http://localhost:15672](http://localhost:15672)

### 5. Verify Installation

Check that all services are healthy:

```bash
# Gateway
curl http://localhost:8080/health

# User Service
curl http://localhost:8080/users/health

# Movie Service
curl http://localhost:8080/movies/health

# Review Service
curl http://localhost:8080/reviews/health
```

## Environment Configuration

### Required Environment Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `POSTGRES_USER` | PostgreSQL username | `alaska_user` |
| `POSTGRES_PASSWORD` | PostgreSQL password | `securepass123` |
| `POSTGRES_DB` | PostgreSQL database name | `alaska_db` |
| `MONGO_USER` | MongoDB username | `mongo_user` |
| `MONGO_PASSWORD` | MongoDB password | `mongopass123` |
| `REDIS_PASSWORD` | Redis password | `redispass123` |
| `RABBITMQ_USER` | RabbitMQ username | `rabbitmq_user` |
| `RABBITMQ_PASSWORD` | RabbitMQ password | `rabbitpass123` |
| `JWT_SECRET` | JWT signing secret (min 32 chars) | `your-very-long-secret-key-here` |
| `OMDB_API_KEY` | OMDb API key | `abc12345` |

### Optional Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `DISCORD_CLIENT_ID` | Discord OAuth client ID | - |
| `DISCORD_CLIENT_SECRET` | Discord OAuth secret | - |
| `GATEWAY_PORT` | Gateway service port | `8080` |
| `USER_SERVICE_PORT` | User service port | `10001` |
| `MOVIE_SERVICE_PORT` | Movie service port | `10002` |
| `REVIEW_SERVICE_PORT` | Review service port | `10003` |

## Development

### Frontend Development (with hot-reload)

```bash
cd frontend
npm install
npm run dev
```

The development server will start at [http://localhost:5173](http://localhost:5173) with hot module replacement enabled.

### Backend Development

To develop a single service locally while using Docker for infrastructure:

```bash
# Start only the infrastructure services
docker-compose up postgres mongodb redis rabbitmq

# Run a service locally
cd user-service
go run cmd/app/main.go
```

### Viewing Logs

```bash
# All services
docker-compose logs -f

# Specific service
docker-compose logs -f user-service
docker-compose logs -f movie-service
docker-compose logs -f review-service
docker-compose logs -f gateway-service
```

### Stopping the Application

```bash
# Stop all services
docker-compose down

# Stop and remove all volumes (WARNING: destroys all data)
docker-compose down -v
```

## Project Structure

```
Alaska/
├── docker-compose.yml           # Orchestration configuration
├── .env                         # Environment variables (gitignored)
├── .gitignore                   # Git ignore rules
├── CLAUDE.md                    # AI assistant documentation
├── README.md                    # This file
│
├── gateway-service/             # API Gateway
│   ├── cmd/app/main.go
│   ├── internal/
│   │   ├── handlers/            # Proxy handlers
│   │   └── utils/               # Config and utilities
│   ├── Dockerfile
│   └── go.mod
│
├── user-service/                # User authentication service
│   ├── cmd/app/main.go
│   ├── internal/
│   │   ├── controllers/         # HTTP handlers
│   │   ├── services/            # Business logic
│   │   ├── repositories/        # Data access
│   │   ├── models/              # Domain models
│   │   ├── dtos/                # Request/Response DTOs
│   │   └── utils/               # Config and utilities
│   ├── Dockerfile
│   └── go.mod
│
├── movie-service/               # Movie data service
│   ├── cmd/app/main.go
│   ├── internal/
│   │   ├── controllers/
│   │   ├── services/
│   │   ├── repositories/
│   │   └── models/
│   ├── Dockerfile
│   └── go.mod
│
├── review-service/              # Review management service
│   ├── cmd/app/main.go
│   ├── internal/
│   │   ├── controllers/
│   │   ├── services/
│   │   ├── repositories/
│   │   └── models/
│   ├── Dockerfile
│   └── go.mod
│
├── shared-events/               # Shared event definitions
│   ├── events/
│   │   ├── user_events.go
│   │   ├── review_events.go
│   │   └── movie_events.go
│   └── go.mod
│
└── frontend/                    # Vue.js frontend
    ├── src/
    │   ├── components/          # Reusable components
    │   ├── views/               # Page components
    │   ├── router/              # Vue Router config
    │   ├── services/            # API services
    │   ├── App.vue
    │   └── main.js
    ├── public/
    ├── package.json
    ├── vite.config.js
    └── Dockerfile
```

## API Documentation

### User Service (`/users`)

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/users/register` | Register a new user | No |
| POST | `/users/login` | Login user | No |
| GET | `/users/me` | Get current user info | Yes |
| POST | `/users/logout` | Logout user | Yes |
| POST | `/users/delete` | Delete user account | Yes |
| GET | `/users/health` | Health check | No |

### Movie Service (`/movies`)

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/movies/:id` | Get movie by IMDb ID | No |
| POST | `/movies` | Create/update movie | No |
| DELETE | `/movies/:id` | Delete movie | No |
| GET | `/movies/health` | Health check | No |

### Review Service (`/reviews`)

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/reviews` | Create a review | Yes |
| GET | `/reviews/movie/:movieId` | Get reviews for a movie | No |
| GET | `/reviews/user/:userId` | Get reviews by a user | No |
| GET | `/reviews/health` | Health check | No |

### Example Requests

**Register User**
```bash
curl -X POST http://localhost:8080/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "moviefan",
    "email": "fan@example.com",
    "password": "securepassword"
  }'
```

**Login**
```bash
curl -X POST http://localhost:8080/users/login \
  -H "Content-Type: application/json" \
  -c cookies.txt \
  -d '{
    "username": "moviefan",
    "password": "securepassword"
  }'
```

**Get Movie**
```bash
curl http://localhost:8080/movies/tt0111161
```

**Create Review**
```bash
curl -X POST http://localhost:8080/reviews \
  -H "Content-Type: application/json" \
  -b cookies.txt \
  -d '{
    "movie_id": "tt0111161",
    "rating": 9,
    "comment": "Excellent movie!"
  }'
```

## Services

### Gateway Service

The API Gateway is the single entry point for all client requests. It handles:
- Request routing to appropriate microservices
- CORS configuration
- Health check aggregation

**Technology**: Go with Chi router

### User Service

Manages user authentication and account operations.

**Features**:
- User registration with email validation
- Secure password hashing (bcrypt)
- Session management with Redis
- Discord OAuth integration
- Account deletion with cascade events

**Technology**: Go, PostgreSQL, Redis, RabbitMQ

### Movie Service

Handles movie data retrieval and caching.

**Features**:
- OMDb API integration
- MongoDB caching for performance
- Movie CRUD operations
- Automatic data updates

**Technology**: Go, MongoDB, RabbitMQ

### Review Service

Manages movie reviews and ratings.

**Features**:
- Review creation and retrieval
- Rating system (1-10)
- User and movie associations
- Cascade deletion on user/movie removal

**Technology**: Go, PostgreSQL, RabbitMQ

## Contributing

We welcome contributions! Here's how you can help:

### Getting Started

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/your-feature-name`
3. Make your changes
4. Run tests (when available)
5. Commit your changes: `git commit -m 'feat: Add some feature'`
6. Push to the branch: `git push origin feature/your-feature-name`
7. Open a Pull Request

### Commit Convention

This project follows [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>(<scope>): <description>

[optional body]

[optional footer]
```

**Types**:
- `feat`: New feature
- `fix`: Bug fix
- `refactor`: Code refactoring
- `build`: Build system changes
- `chore`: Maintenance tasks
- `docs`: Documentation updates

**Scopes**: `Frontend`, `User`, `Movie`, `Review`, `Gateway`, `Docker`, `Infra`

**Examples**:
- `feat(Frontend): Add movie search functionality`
- `fix(User): Resolve session expiration bug`
- `refactor(Review): Improve review query performance`

### Code Style

- **Go**: Follow standard Go conventions, run `go fmt`
- **JavaScript/Vue**: Follow Vue.js style guide
- **Comments**: Write clear, concise comments for complex logic
- **Documentation**: Update CLAUDE.md for significant architectural changes

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [OMDb API](http://www.omdbapi.com/) for movie data
- [Ollama](https://ollama.ai/) for local LLM integration
- All open-source libraries and tools used in this project

## Support

For issues, questions, or contributions, please:
- Open an issue on GitHub
- Check existing documentation in CLAUDE.md
- Review closed issues for similar problems

---

**Built with love by the Alaska team**
