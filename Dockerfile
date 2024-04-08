FROM golang:1.21.6-alpine AS builder

RUN apk add --no-cache \
    git \
    gcc \
    libc-dev \
    curl \
    openssl

WORKDIR /opt/user-manager
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /usr/bin/user-manager .

EXPOSE 3000

CMD ["/usr/bin/user-manager"]



