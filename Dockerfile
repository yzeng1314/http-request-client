## Build
FROM --platform=$BUILDPLATFORM golang:1.18-alpine AS build

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

ARG TARGETOS TARGETARCH
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /http-client

## Deploy
FROM alpine:latest

WORKDIR /

COPY --from=build /http-client /http-client 

RUN addgroup -S requsr \
    && adduser -S requsr -G requsr \
    && apk --no-cache add curl

USER requsr

ENTRYPOINT [ "/http-client" ]