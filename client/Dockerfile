WORKDIR /app
COPY ./main.go .
COPY ../pkg ./pkg
RUN go build -o client main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/client .
CMD ["./client"]
