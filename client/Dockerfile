# Etapa de build
FROM golang:1.21 AS builder
WORKDIR /app
COPY ./client/main.go .
COPY ./pkg ./pkg
RUN go build -o client main.go

# Etapa de execução
FROM alpine
WORKDIR /app
COPY --from=builder /app/client .
CMD ["./client"]
