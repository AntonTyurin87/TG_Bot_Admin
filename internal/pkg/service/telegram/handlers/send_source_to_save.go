package telegram

import (
	"TG_Bot_Admin/internal/pkg/domain/entity"
	"TG_Bot_Admin/internal/pkg/service/telegram/auth"
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (h *Handler) sendSourceToSaveHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})

	userID := update.CallbackQuery.From.ID
	chatID := update.CallbackQuery.Message.Message.Chat.ID

	if auth.SuperAdmin != auth.GetUserCategory(userID) {
		// отправляемся в меню по умолчанию
		h.DefaultAnswerMenu(ctx, b, chatID, librarian_admin_start)
	}

	if h.adminService.IsAnyNotFinishedSource(ctx, userID) && h.adminService.IsNowStep(ctx, entity.SourceDownloadURLStep, userID) {
		// записываем текущую дату и ставим следующий шаг
		h.adminService.UpdateLibrarianSourceItem(ctx, userID, "")

		h.createSourceDefaultMenu(ctx, b, chatID, userID)
	}

}
