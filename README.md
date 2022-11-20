# Chat Server

## App Environment Variable

  Recommand to make `.env` file for env variable to reuse in docker config
  - Mysql : `MYSQL_USER`, `MYSQL_PASSWORD`, `MYSQL_HOST`, `MYSQL_PORT`, `MYSQL_DATABASE`
  - Redis : `REDIS_HOST`, `REDIS_PASSWORD`, `REDIS_PORT`
  - Server : `REST_PORT`, `GRPC_PORT`

## Development

  1. install air

  📝 configure `$GOBIN` path before installing 

    go install github.com/cosmtrek/air@latest

  2. run app with air

    air
