# Этап сборки
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o todo_list ./cmd/web-app

FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/todo_list /app/todo_list

COPY --from=builder /app/internal/database/postgresql/migrations /app/migrations

EXPOSE 3000

CMD ["/app/todo_list"]
