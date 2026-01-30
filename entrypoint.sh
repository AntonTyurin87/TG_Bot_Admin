#!/bin/bash
set -e

echo "=== Запуск entrypoint ==="

# Проверка переменных
if [ -z "$TG_BOT_ADMIN_TOKEN" ]; then
    echo "❌ ОШИБКА: TG_BOT_ADMIN_TOKEN не установлен!"
    echo "Переменные окружения в контейнере:"
    env | grep TG_ || echo "Нет TG_ переменных"
    exit 1
fi

echo "✅ TG_BOT_ADMIN_TOKEN установлен (длина: ${#TG_BOT_ADMIN_TOKEN})"

# Запуск приложения
exec "$@"