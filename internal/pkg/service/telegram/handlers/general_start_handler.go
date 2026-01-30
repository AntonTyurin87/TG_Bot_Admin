package telegram

import (
	"TG_Bot_Admin/internal/pkg/service/telegram/auth"
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (h *Handler) generalStartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})

	userID := update.CallbackQuery.From.ID
	chatID := update.CallbackQuery.Message.Message.Chat.ID

	// проверка прав пользователя
	userCategory := auth.GetUserCategory(userID)

	switch {
	case userCategory == auth.SuperAdmin:
		h.generalStartSuperAdminMenu(ctx, b, chatID)
	case userCategory == auth.LibrarianAdmin:
		h.generalStartLibrarianAdminMenu(ctx, b, chatID)
	default:
		h.generalStartSimpleUserMenu(ctx, b, chatID)
	}

}

func (h *Handler) generalStartSuperAdminMenu(ctx context.Context, b *bot.Bot, chatID int64) {
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: Library, CallbackData: super_admin_library},
			},
			{
				{Text: Empty, CallbackData: default_menu}, //TODO сделать handler
			},
			{
				{Text: BackTo + ReconComGroup, URL: "https://t.me/+qbEymR_JfXFhOWUy"}, //TODO пересобрать на адрес из переменной окружения
			},
		},
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        SuperAdminStartMenu,
		ReplyMarkup: kb,
		ParseMode:   models.ParseModeMarkdown,
	})

}

func (h *Handler) generalStartLibrarianAdminMenu(ctx context.Context, b *bot.Bot, chatID int64) {
	h.DefaultAnswerMenu(ctx, b, chatID, librarian_admin_start) //TODO меню для библиотекаря
}

func (h *Handler) generalStartSimpleUserMenu(ctx context.Context, b *bot.Bot, chatID int64) {
	h.DefaultAnswerMenu(ctx, b, chatID, simple_user_start) //TODO меню для простого пользователя
}
