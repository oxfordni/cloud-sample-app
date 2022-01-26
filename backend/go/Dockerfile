FROM golang:1.17-alpine AS worker

WORKDIR /app

# We need curl to be installed
RUN apk --no-cache add curl

# Install live-reloading for Go applications in development (air)
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# Make sure 'air' is in the PATH
RUN echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> /etc/profile

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 3000
