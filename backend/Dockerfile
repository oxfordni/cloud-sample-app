FROM golang:1.16.5-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go get github.com/pilu/fresh

RUN go mod download

COPY . .

EXPOSE 3000
