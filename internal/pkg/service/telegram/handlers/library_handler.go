package telegram

import (
	"context"
	"fmt"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// libraryCallbackHandler - Обработчик callback для кнопки "Что есть в библиотеке"
func libraryCallbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})

	sendLibraryCallbackHandler(ctx, b, update.CallbackQuery.Message.Message.Chat.ID)
}

func sendLibraryCallbackHandler(ctx context.Context, b *bot.Bot, chatID int64) {
	text := fmt.Sprint(
		"📁 *Что есть в библиотеке*\n\nВ этом разделе будет приведена статистика того, что сейчас имеется в библиотеке\\.",
	)

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    chatID,
		Text:      text,
		ParseMode: models.ParseModeMarkdown,
		ReplyMarkup: &models.InlineKeyboardMarkup{
			InlineKeyboard: [][]models.InlineKeyboardButton{
				{
					{Text: "🗂 Новые источники за 7 дней", CallbackData: "new_sources"},
				},
				{
					{Text: "🔙 Назад в главное меню", CallbackData: "menu"},
				},
			},
		},
	})
	if err != nil {
		log.Printf("Error sending search menu: %v", err)
	}
}
