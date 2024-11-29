# Этап 1: Сборка
FROM golang:latest AS builder
WORKDIR /app
ADD go.mod .
COPY . .
RUN go build -o main main.go

# Этап 2: Минимальный финальный образ
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main /app/main
CMD ["./main"]
