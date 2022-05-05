FROM golang:1.14-alpine AS build

WORKDIR /app

COPY go.mod .
 
COPY go.sum .

RUN go mod download 

COPY . .

RUN go build

CMD ["/app/main"]
