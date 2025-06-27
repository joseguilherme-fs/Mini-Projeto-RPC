FROM golang:1.21 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./client/main.go ./client/main.go
COPY ./pkg ./pkg

RUN go build -o client ./client/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/client .
CMD ["./client"]
