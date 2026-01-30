package telegram

import (
	"TG_Bot_Admin/internal/pkg/service/telegram/helpers"
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// DefaultAnswerMenu ...
func (h *Handler) DefaultAnswerMenu(ctx context.Context, b *bot.Bot, chatID int64, menuName string) {
	messageText := helpers.EscapeMarkdown(fmt.Sprintf("🚧*%s*🚧\n находится в разработке", menuName))

	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "\t✍ Вернуться к началу", CallbackData: general_start},
			},
			{
				{Text: "🔙 Вернуться в Recom_Com", URL: "https://t.me/+qbEymR_JfXFhOWUy"}, //TODO пересобрать на адрес из переменной окружения
			},
		},
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        messageText,
		ParseMode:   models.ParseModeMarkdown,
		ReplyMarkup: kb,
	})
}
