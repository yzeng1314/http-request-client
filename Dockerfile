## Build
FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /http-client

## Deploy
FROM alpine:latest

WORKDIR /

COPY --from=build /http-client /http-client 

ENTRYPOINT [ "/http-client" ]