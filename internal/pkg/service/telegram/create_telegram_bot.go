package telegram

import (
	"TG_Bot_Admin/internal/pkg/service/telegram/auth"
	handlers "TG_Bot_Admin/internal/pkg/service/telegram/handlers"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var TargetChatID int64 = -1002652107754 // Значение по умолчанию
var TargetThreadID int = 256            // Значение по умолчанию

//// CreateTelegramBot создает и настраивает Telegram бота
//func CreateTelegramBot(handler *handlers.Handler) (*bot.Bot, error) {
//	opts := []bot.Option{
//		bot.WithDefaultHandler(handlers.DefaultHandler),
//	}
//
//	botToken := os.Getenv("TG_BOT_ADMIN_TOKEN") //TODO заменить на константу
//
//	b, err := bot.New(botToken, opts...)
//	if err != nil {
//		return nil, err
//	}
//
//	// Регистрация обработчиков
//	handler.RegisterHandlers(b)
//
//	return b, nil
//}

const (
	GROUP_ID   = int64(-1002652107754)
	TOPIC_ID   = 256
	TOPIC_LINK = "https://web.telegram.org/a/#-1002652107754_256"
)

// CreateTelegramBot создает и настраивает Telegram бота
func CreateTelegramBot(handler *handlers.Handler) (*bot.Bot, error) {
	// Создаем кастомный HTTP клиент с увеличенными таймаутами
	httpClient := &http.Client{
		Timeout: 120 * time.Second, // Таймаут для отдельных запросов
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     90 * time.Second,
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}

	opts := []bot.Option{
		bot.WithDefaultHandler(handlers.DefaultHandler),
		bot.WithHTTPClient(55*time.Second, httpClient),
		bot.WithCheckInitTimeout(10 * time.Second),
		bot.WithServerURL("https://api.telegram.org"), // Явно указываем URL
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

func registerHandlers(b *bot.Bot) {
	// 1. Приветствие в топике
	b.RegisterHandler(bot.HandlerTypeMessageText, "/admin_topic_start", bot.MatchTypeExact,
		func(ctx context.Context, b *bot.Bot, update *models.Update) {
			// Отправляем в топик приглашение
			kb := &models.InlineKeyboardMarkup{
				InlineKeyboard: [][]models.InlineKeyboardButton{
					{{Text: "🤖 Начать диалог",
						URL: fmt.Sprintf("https://t.me/ReconV1_Bot?start=from_topic_%d", TOPIC_ID)}},
				},
			}

			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:          GROUP_ID,
				MessageThreadID: TOPIC_ID,
				Text:            "👋 Напишите боту /start в личные сообщения для индивидуальной работы!",
				ReplyMarkup:     kb,
			})
		})

	// 2. /start в личных сообщениях
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact,
		func(ctx context.Context, b *bot.Bot, update *models.Update) {
			user := update.Message.From

			userCategory := auth.GetUserCategory(user.ID)

			switch {
			case userCategory == auth.SuperAdmin:
				// Меню с 1 кнопкой
				kb := &models.InlineKeyboardMarkup{
					InlineKeyboard: [][]models.InlineKeyboardButton{
						{{Text: fmt.Sprintf("👤 Меню для Супер Администратора!"),
							CallbackData: "personal_menu"}},
					},
				}

				b.SendMessage(ctx, &bot.SendMessageParams{
					ChatID:      update.Message.Chat.ID,
					Text:        fmt.Sprintf("Привет из топика! 👋 Я буду работать с тобой, %s", user.FirstName),
					ReplyMarkup: kb,
				})
			default:
				// Меню с 1 кнопкой
				b.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: update.Message.Chat.ID,
					Text:   "👋 Напишите боту /start в личные сообщения для индивидуальной работы!",
				})

			}

		})

	// 3. Личное меню
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "personal_menu", bot.MatchTypeExact,
		func(ctx context.Context, b *bot.Bot, update *models.Update) {
			callback := update.CallbackQuery
			user := callback.From

			b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
				CallbackQueryID: callback.ID,
			})

			b.DeleteMessage(ctx, &bot.DeleteMessageParams{
				ChatID:    callback.Message.Message.Chat.ID,
				MessageID: callback.Message.Message.ID,
			})

			// Меню с 2 кнопками
			kb := &models.InlineKeyboardMarkup{
				InlineKeyboard: [][]models.InlineKeyboardButton{
					{{Text: "ℹ️ Моя информация", CallbackData: "my_info"}},
					{{Text: "💬 Написать", CallbackData: "send_message"}},
				},
			}

			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:      callback.Message.Message.Chat.ID,
				Text:        fmt.Sprintf("👤 *Личное меню для %s*\n\nВыберите:", user.FirstName),
				ParseMode:   models.ParseModeMarkdown,
				ReplyMarkup: kb,
			})
		})

	// 4. Кнопка "Моя информация"
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "my_info", bot.MatchTypeExact,
		func(ctx context.Context, b *bot.Bot, update *models.Update) {
			callback := update.CallbackQuery

			b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
				CallbackQueryID: callback.ID,
			})

			info := fmt.Sprintf(
				"👤 *Ваша информация*\n\n"+
					"🆔 ID: `%d`\n"+
					"👤 Имя: %s\n"+
					"🔗 Username: @%s\n\n"+
					"Вы подключились из топика:\n%s",
				callback.From.ID,
				callback.From.FirstName,
				callback.From.Username,
				TOPIC_LINK,
			)

			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:    callback.Message.Message.Chat.ID,
				Text:      info,
				ParseMode: models.ParseModeMarkdown,
			})
		})

	// 5. Кнопка "Написать"
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "send_message", bot.MatchTypeExact,
		func(ctx context.Context, b *bot.Bot, update *models.Update) {
			callback := update.CallbackQuery

			b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
				CallbackQueryID: callback.ID,
			})

			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: callback.Message.Message.Chat.ID,
				Text: fmt.Sprintf("%s, напишите что-нибудь, и я отвечу индивидуально! ✨",
					callback.From.FirstName),
			})
		})

	// 6. Ответы на текстовые сообщения
	b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypePrefix,
		func(ctx context.Context, b *bot.Bot, update *models.Update) {
			if update.Message.Text == "/start" {
				return
			}

			user := update.Message.From
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   fmt.Sprintf("%s, это наш индивидуальный диалог! ✨", user.FirstName),
			})
		})
}
