FROM golang:1.16.5-alpine

WORKDIR /app

COPY . .

RUN go get github.com/pilu/fresh

RUN go mod download

EXPOSE 3000
