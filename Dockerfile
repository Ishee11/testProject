# Этап 1: Сборка
FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Этап 2: Минимальный финальный образ
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
RUN apk add --no-cache libc6-compat
EXPOSE 8080
CMD ["./main"]
