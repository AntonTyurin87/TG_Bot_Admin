package telegram

import (
	handlers "TG_Bot_Admin/internal/pkg/service/telegram/handlers"
	"os"

	"github.com/go-telegram/bot"
)

// CreateTelegramBot создает и настраивает Telegram бота
func CreateTelegramBot(handler *handlers.Handler) (*bot.Bot, error) {
	opts := []bot.Option{
		bot.WithDefaultHandler(handlers.DefaultHandler),
	}

	botToken := os.Getenv("TG_BOT_ADMIN_TOKEN") //TODO заменить на константу

	b, err := bot.New(botToken, opts...)
	if err != nil {
		return nil, err
	}

	// Регистрация обработчиков
	handler.RegisterHandlers(b)

	return b, nil
}
