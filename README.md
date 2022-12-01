# Chat Server

## Structure
  **1. API Service**
  
  - Environment Variable  
    - Mysql : `MYSQL_USER`, `MYSQL_PASSWORD`, `MYSQL_HOST`, `MYSQL_PORT`, `MYSQL_DATABASE`
    - Redis : `REDIS_HOST`, `REDIS_PASSWORD`, `REDIS_PORT`
    - Server : `REST_PORT`, `GRPC_PORT`

  **2. Chat Service**

  **3. Media Service**

## Development

  1. run docker
  2. install air(live-reload)

  📝 configure `$GOBIN` path before installing 

    go install github.com/cosmtrek/air@latest

  3. run app with air

    air
