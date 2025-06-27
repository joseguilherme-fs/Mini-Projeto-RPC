FROM golang:1.21 AS builder
WORKDIR /app
COPY ./server/main.go .
COPY ./pkg ./pkg
RUN go build -o server main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/server .
CMD ["./server"]
