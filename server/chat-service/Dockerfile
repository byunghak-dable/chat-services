# build
FROM golang:1.22.0-alpine3.19  AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build cmd/main.go 

# deploy
FROM alpine:3.16

WORKDIR /

COPY --from=builder /app/main /main

ENTRYPOINT ["main"]
