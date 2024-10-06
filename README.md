# DevSeConnect 🌐

Welcome to **DevSeConnect** – A social media platform for developers to share knowledge and collaborate on DevOps and software development topics.

## Project Overview 🛠️

This repository contains the backend and frontend implementation of DevSeConnect, built using **Go (Golang) and Svelte**. The project utilizes several modern technologies including **RabbitMQ**, **PostgreSQL**, **Redis**, **Svelte** and **Docker** to provide a scalable and performant architecture.

## Folder Structure 📁

```plaintext
├── cmd
│   ├── main.go      # Main entry point of the application           
├── internal
│   ├── application
│   │   └── service
│   │       ├── comment_service.go
│   │       ├── like_service.go
│   │       ├── message_service.go
│   │       ├── post_service.go
│   │       ├── post_tags_service.go
│   │       ├── tags_service.go
│   │       └── user_service.go
│   ├── domain
│   │   ├── entity
│   │   │   ├── comments.go
│   │   │   ├── likes.go
│   │   │   ├── messages.go
│   │   │   ├── network.go
│   │   │   ├── posts.go
│   │   │   ├── posttags.go
│   │   │   ├── tags.go
│   │   │   ├── user_connections.go
│   │   │   └── users.go
│   │   ├── handler
│   │   │   ├── comment_handler.go
│   │   │   ├── like_handler.go
│   │   │   ├── message_handler.go
│   │   │   ├── post_handler.go
│   │   │   ├── post_tags_handler.go
│   │   │   ├── tag_handler.go
│   │   │   └── user_handler.go
│   │   └── repository
│   │       ├── comment_repository.go
│   │       ├── likes_repository.go
│   │       ├── message_repository.go
│   │       ├── post_repository.go
│   │       ├── post_tags_repository.go
│   │       ├── tags_repository.go
│   │       └── user_repository.go
│   └── infrastructure
│       ├── cache
│       │   └── redis.go
│       ├── database
│       │   └── postgres.go
│       ├── messaging
│       │   ├── comment_consumer.go
│       │   ├── consumer.go
│       │   ├── like_consumer.go
│       │   ├── message_consumer.go
│       │   ├── post_consumer.go
│       │   ├── post_tags_consumer.go
│       │   ├── producer.go
│       │   ├── rabbitmq.go
│       │   ├── tags_consumer.go
│       │   └── user_consumer.go
│       └── routing
│           └── router.go
├── pkg
│   ├── response
│   │   └── error.go
│   └── security
│       └── hash.go
├── scripts
│   ├── migrate.sh
│   └── migrations
│       ├── devseconnect_insert.sql
│       ├── devseconnect.sql
│       ├── devseconnect_test_select2.sql
│       └── devseconnect_test_select.sql
├── .gitignore                     # Git ignore rules
├── docker-compose.yml             # Docker Compose setup
├── Dockerfile                     # Dockerfile for containerization
├── nginx.conf
├── nohup.out
├── go.mod                         # Go module dependencies
├── prometheus.yml
├── LICENSE                        # License file
```
## Technologies Used 🛠️

- **Go (Golang)**
- **Svelte** - Web Framework.
- **PostgreSQL** – For database management.
- **Redis** – Caching layer.
- **RabbitMQ** – Message broker for async communication.
- **Docker** – Containerization of the application.
- **Grafana** – Monitoring graph visualization.
- **Prometheus** – Monitoring tool.

## Installation 🚀

### Prerequisites

Make sure you have the following installed:

- **Go v1.18+**
- **Svelte 5**
- **Docker** & **Docker Compose**
- **Redis**
- **RabbitMQ**
- **PostgreSQL**
- **Prometheus**
- **Grafana**

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
   docker-compose up -d --build
   ```

4. **Run locally** (without Docker):
   ```bash
   go run cmd/main.go
   ```
   
5. **Log Webserver**:
   ```bash
   docker logs -f golang_web_server
   ```

## API Endpoints 🔗

The platform uses RESTful APIs to handle different operations:

- **Users**:

  - `GET /users` – Fetch all users.
  - `POST /users` – Create a new user.
  - `PUT /users/:id` – Update a user with id.
  - `DELETE /users/:id` – Delete a user with id.

- **Posts**:

  - `GET /posts` – Fetch all posts.
  - `POST /posts/:username` – Create a new post.
  - `PUT /posts/:id` – Update a post with id.
  - `DELETE /posts/:id` – Delete a post with id.

- **Comments**:

  - `GET /comments` – Fetch all comments.
  - `POST /comments/:title/:username` – Add a comment to a post.
  - `PUT /comments/:id` – Update a comment with id.
  - `DELETE /comments/:id` – Delete a comment with id.

- **Likes**:
  - `GET /likes` – Fetch all likes.
  - `POST /likes/:title/:username` – Like a post.
  - `POST /likes/:content/:username` – Like a comment.
  - `DELETE /likes/:id` – Delete a like with id.

## Database Schema 🗄️

The database is managed using **PostgreSQL**, and migration scripts are located in `scripts/migrations/`. The key tables include:

- `users`
- `posts`
- `comments`
- `messages`
- `likes`
- `tags`
- `post_tags`

## Messaging System ✉️

Asynchronous communication between services is managed via **RabbitMQ**. Producers and consumers are defined in the `internal/infrastructure/messaging/` directory.

## Monitoring 📊

**Prometheus** is used for application monitoring and metrics collection. The configuration is located in `prometheus.yml`.

## License 📜

This project is licensed under the [MIT License](./LICENSE).
