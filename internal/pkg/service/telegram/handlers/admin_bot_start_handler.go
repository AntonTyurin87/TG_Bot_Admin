package telegram

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (h *Handler) adminBotStartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	h.sendAdminStartMenu(ctx, b, update.Message.Chat.ID)
}

func (h *Handler) sendAdminStartMenu(ctx context.Context, b *bot.Bot, chatID int64) {
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "\t✍ Начать работу", CallbackData: general_start},
			},
			{
				{Text: "🔙 Вернуться в Recom_Com", URL: "https://t.me/+qbEymR_JfXFhOWUy"}, //TODO пересобрать на адрес из переменной окружения
			},
		},
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        "Это Бот-Администратор и тут какой-то текст про этого бота!", //TODO сделать текст про этого бота
		ReplyMarkup: kb,
	})
}
