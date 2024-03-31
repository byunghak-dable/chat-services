# Chat App

This is a toy project to use **Kafka** & **API Gateway**, and the server app is composed of Hexagonal Architecture.

## ✨ Features

- 🦼 **User Server**
  - Manage user authentication using Oauth2(Google, Github) and JWT
- 💬 **Chat Server**
  - Chat server using web socket
  - Broadcast messages to chat members participating in
    different server instances through kafka message broker
- 🚪 **Api Gateway(TODO)**

## ⚙️ Requirements

- Docker >= 26.0.0
- Java 17 (for **User Service**)
- Go 1.22 (for **Chat Service**)

## 🚀 Getting Started

### User Service

- **Set Environment**

  - **Server** - `REST_PORT`
  - **MongoDB** - `MONGO_HOST`, `MONGO_PORT`, `MONGO_USER`, `MONGO_PASSWORD`, `MONGO_DB`, `MONGO_AUTH_DB`
  - **Oauth2(Google)** - `GOOGLE_OAUTH_CLIENT_ID`, `GOOGLE_OAUTH_CLIENT_SECRET`, `GOOGLE_OAUTH_REDIRECT_URI`
  - **Oauth2(Github)** - `GITHUB_OAUTH_CLIENT_ID`, `GITHUB_OAUTH_CLIENT_SECRET`, `GITHUB_OAUTH_REDIRECT_URI`

- **Run Docker**

  ```sh
  docker-compose up -d
  ```

- **Run Server**

  ```sh
  ${USER_SERVICE_ROOT}/run.sh
  ```

### Chat Service

- **Set Environment**

  - **Server** - `REST_PORT`
  - **MongoDB** - `MONGO_HOST`, `MONGO_PORT`, `MONGO_USER`, `MONGO_PASSWORD`, `MONGO_DB`, `MONGO_AUTH_DB`
  - **Kafka** - `KAFKA_KRAFT_CLUSTER_ID`, `KAFKA_{1..3}_HOST`, `KAFKA_{1..3}_PORT`, `KAFKA_CHAT_TOPIC`, `KAFKA_CLIENT_ID`, `KAFKA_GROUP_ID`

- **Run Docker**

  ```sh
  docker-compose up -d
  ```

- **Run Server**

  ```sh
  go run ${CHAT_SERVICE_ROOT}/cmd/main/main.go

  ```
