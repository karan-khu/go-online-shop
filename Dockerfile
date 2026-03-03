FROM golang:1.25.0-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o main main.go

EXPOSE 8800

CMD ["./main"]
