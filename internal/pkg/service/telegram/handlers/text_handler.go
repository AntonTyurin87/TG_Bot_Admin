package telegram

import (
	"context"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Обработчик текстовых сообщений
func (h *Handler) messageHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	userID := update.Message.From.ID
	text := update.Message.Text

	// Проверяем, не является ли сообщение командой или файлом
	if strings.HasPrefix(text, "/") {
		return
	}

	// Проверяем состояние пользователя
	if state, ok := UserSourceStates[userID]; ok && state == CreateSource {
		if update.Message.Document != nil {
			h.fileMessageHandler(ctx, b, update)
			return
		}

		// записываем текс в БД
		h.adminService.UpdateLibrarianSourceItem(ctx, userID, text)
		// вызываем дефолтный хендлер для сохранения источника
		h.createSourceDefaultHandler(ctx, b, update)
	}
}
