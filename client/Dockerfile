FROM golang:1.21 AS builder
WORKDIR /app
COPY ../pkg ./pkg
COPY ./main.go .
RUN go build -o client main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/client .
CMD ["./client"]
