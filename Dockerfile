FROM golang:1.25-alpine AS builder

WORKDIR /app

# Копируем и устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN go build -o tg_bot_admin ./cmd/tg_bot_admin/main.go

# Финальный образ
FROM alpine:latest

RUN apk --no-cache add ca-certificates bash

WORKDIR /root/

# Копируем бинарник из builder stage
COPY --from=builder /app/tg_bot_admin .

# Копируем entrypoint.sh из текущей директории
COPY entrypoint.sh .

# Делаем скрипт исполняемым
RUN chmod +x entrypoint.sh

# Проверяем, что файлы на месте
RUN ls -la

# Точка входа
ENTRYPOINT ["./entrypoint.sh"]
CMD ["./tg_bot_admin"]