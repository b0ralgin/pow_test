FROM golang:1.23-alpine AS client

RUN apk add --no-cache make

WORKDIR /app

COPY . .

RUN make build-client

ENTRYPOINT ["./bin/client"]

FROM golang:1.23-alpine AS server

RUN apk add --no-cache make

WORKDIR /app

COPY . .

RUN make build-server

ENTRYPOINT ["./bin/server"]