package telegram

import (
	"TG_Bot_Admin/internal/pkg/service/telegram/auth"
	"fmt"

	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (h *Handler) createSourceDefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	userID := update.Message.From.ID
	chatID := update.Message.Chat.ID

	if auth.SuperAdmin != auth.GetUserCategory(userID) {
		// отправляемся в меню по умолчанию
		h.DefaultAnswerMenu(ctx, b, chatID, librarian_admin_start)
	}

	h.createSourceDefaultMenu(ctx, b, chatID, userID)
}

func (h *Handler) createSourceDefaultMenu(ctx context.Context, b *bot.Bot, chatID, userID int64) {
	source, err := h.adminService.SelectLibrarianSourceItem(ctx, userID)
	if err != nil {
		fmt.Errorf("h.adminService.SelectLibrarianSourceItem: %w", err)
	}

	if source == nil {
		return
	}

	// получаем текст для заголовка сообщения
	messageText := h.presenter.TextMessageToCreateSource(source)

	kb := h.presenter.KeyBlockToCreateSource(source)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        messageText,
		ReplyMarkup: kb,
		ParseMode:   models.ParseModeMarkdown,
	})
}
