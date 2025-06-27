FROM golang:1.21 AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY ./server/main.go ./server/main.go
COPY ./pkg ./pkg

RUN go build -o server ./server/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/server .
CMD ["./server"]
