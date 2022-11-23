# build
FROM golang:1.19.3-alpine3.16  AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build cmd/main.go 

# deploy
FROM alpine:3.16

WORKDIR /

COPY --from=build /app/main /main

ENTRYPOINT ["main"]
