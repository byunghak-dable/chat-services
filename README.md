# Chat Server

## App Environment Variable

  Recommand to make `.env` file for env variable to reuse in docker config

  - `MYSQL_USER`
  - `MYSQL_PASSWORD`
  - `MYSQL_HOST`
  - `MYSQL_PORT`
  - `MYSQL_DATABASE`
  - `REST_PORT`

## Development

  1. install air

  📝 configure `$GOBIN` path before installing 

    ```bash
    go install github.com/cosmtrek/air@latest
    ``` 
  2. run app with air

    ```bash
    air
    ``` 
