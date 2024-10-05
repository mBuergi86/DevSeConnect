# DevSeConnect 🌐

Welcome to **DevSeConnect** – A social media platform for developers to share knowledge and collaborate on DevOps and software development topics.

## Project Overview 🛠️

This repository contains the backend implementation of DevSeConnect, built using **Go (Golang)**. The project utilizes several modern technologies including **RabbitMQ**, **PostgreSQL**, **Redis**, and **Docker** to provide a scalable and performant architecture.

## Folder Structure 📁

```plaintext
├── cmd/
│   └── main.go                   # Main entry point of the application
├── internal/
│   ├── application/
│   │   └── service/
│   │      ├── comment_service.go
│   │      ├── like_service.go
│   │      ├── message_service.go
│   │      ├── post_service.go
│   │      ├── post_tags_service.go
│   │      ├── tags_service.go
│   │      └── user_service.go
│   ├── domain/
│   │   ├── entity/
│   │   ├── handler/
│   │   └── repository/
│   ├── infrastructure/
│   │   ├── cache/
│   │   │   └── redis.go
│   │   ├── database/
│   │   │   └── postgres.go
│   │   ├── messaging/
│   │   │   ├── producer.go
│   │   │   ├── rabbitmq.go
│   │   │   └── consumers/
│   │   ├── routing/
│   │   │   └── router.go
├── scripts/
│   └── migrations/
├── .gitignore                     # Git ignore rules
├── docker-compose.yml             # Docker Compose setup
├── Dockerfile                     # Dockerfile for containerization
├── go.mod                         # Go module dependencies
├── LICENSE                        # License file
```
## Technologies Used 🛠️

- **Go (Golang)**
- **PostgreSQL** – For database management.
- **Redis** – Caching layer.
- **RabbitMQ** – Message broker for async communication.
- **Docker** – Containerization of the application.
- **Prometheus** – Monitoring tool.

## Installation 🚀

### Prerequisites

Make sure you have the following installed:

- **Go v1.18+**
- **Docker** & **Docker Compose**
- **Redis**
- **RabbitMQ**
- **PostgreSQL**

### Steps to Install

1. **Clone the repository**:

   ```bash
   git clone https://github.com/mBuergi86/devseconnect.git
   cd devseconnect
   ```

2. **Install dependencies**:

   ```bash
   go mod download
   ```

3. **Run with Docker**:

   ```bash
   docker-compose up
   ```

4. **Run locally** (without Docker):
   ```bash
   go run cmd/main.go
   ```

## API Endpoints 🔗

The platform uses RESTful APIs to handle different operations:

- **Users**:

  - `GET /users` – Fetch all users.
  - `POST /users` – Create a new user.

- **Posts**:

  - `GET /posts` – Fetch all posts.
  - `POST /posts/:username` – Create a new post.

- **Comments**:

  - `GET /comments` – Fetch all comments.
  - `POST /comments/:title/:username` – Add a comment to a post.

- **Likes**:
  - `POST /posts/:title/like/:username` – Like a post.
  - `POST /comments/:content/like/:username` – Like a comment.

## Database Schema 🗄️

The database is managed using **PostgreSQL**, and migration scripts are located in `scripts/migrations/`. The key tables include:

- `users`
- `posts`
- `comments`
- `likes`
- `tags`
- `post_tags`

## Messaging System ✉️

Asynchronous communication between services is managed via **RabbitMQ**. Producers and consumers are defined in the `internal/infrastructure/messaging/` directory.

## Monitoring 📊

**Prometheus** is used for application monitoring and metrics collection. The configuration is located in `prometheus.yml`.

## License 📜

This project is licensed under the [MIT License](./LICENSE).
