package telegram

import (
	"TG_Bot_Admin/internal/pkg/domain/entity"
	"TG_Bot_Admin/internal/pkg/domain/texts"
	"TG_Bot_Admin/internal/pkg/service/telegram/auth"
	"TG_Bot_Admin/internal/pkg/service/telegram/helpers"
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (h *Handler) deleteSourceDefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})

	userID := update.CallbackQuery.From.ID
	chatID := update.CallbackQuery.Message.Message.Chat.ID

	if auth.SuperAdmin != auth.GetUserCategory(userID) {
		// отправляемся в меню по умолчанию
		h.DefaultAnswerMenu(ctx, b, chatID, librarian_admin_start)
	}

	h.deleteSourceDefaultMenu(ctx, b, chatID, userID)
}

func (h *Handler) deleteSourceDefaultMenu(ctx context.Context, b *bot.Bot, chatID, userID int64) {
	// проверяем, что есть маркер создания источника пользователем
	if state, ok := UserSourceStates[userID]; !ok && state == CreateSource {
		return //TODO логирование
	}

	// проверяем наличие источника и файлов к нему
	source, err := h.adminService.SelectLibrarianSourceItem(ctx, userID)
	if err != nil {
		return //TODO логирование
	}
	// если источника нет или он готов к отправке, то не удаляем
	if source == nil || source.Step > entity.SourceLoadFileStep {
		return //TODO логирование
	}

	// удаляем источник и файл к нему
	sourceID := h.adminService.DeleteLibrarianSourceItem(ctx, userID)
	if sourceID == 0 {
		_ = h.adminService.DeleteLibrarianSourceFile(ctx, sourceID)
		//TODO логирование
	}

	//TODO формируем набор кнопок - убрать в функцию презентера

	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: BackTo + Library, CallbackData: general_start},
			},
			{
				{Text: BackTo + ReconComGroup, URL: texts.KeyURLReconComGroupURL.String()},
			},
		},
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        helpers.EscapeMarkdown(texts.InstructionsDeleteSourceSuccess.String()),
		ReplyMarkup: kb,
		ParseMode:   models.ParseModeMarkdown,
	})
}
