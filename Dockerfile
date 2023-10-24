FROM golang:latest AS build-env
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.mod /src/
RUN go mod download
COPY . .
RUN go build -a -o gohttp -trimpath

FROM alpine:latest

RUN apk add --no-cache ca-certificates \
    && rm -rf /var/cache/*

RUN mkdir -p /app /mount \
    && adduser -D gohttp \
    && chown -R gohttp:gohttp /app /mount

USER gohttp
WORKDIR /mount

COPY --from=build-env /src/gohttp /app

ENTRYPOINT [ "/app/gohttp" ]