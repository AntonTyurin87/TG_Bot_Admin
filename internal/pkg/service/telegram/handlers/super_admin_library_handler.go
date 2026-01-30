package telegram

import (
	"TG_Bot_Admin/internal/pkg/service/telegram/auth"
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (h *Handler) superAdminLibraryHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})

	userID := update.CallbackQuery.From.ID
	chatID := update.CallbackQuery.Message.Message.Chat.ID

	if auth.SuperAdmin != auth.GetUserCategory(userID) {
		// отправляемся в меню по умолчанию
		h.DefaultAnswerMenu(ctx, b, chatID, librarian_admin_start)
	}

	h.superAdminLibraryMenu(ctx, b, chatID)
}

func (h *Handler) superAdminLibraryMenu(ctx context.Context, b *bot.Bot, chatID int64) {
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: CreateSource, CallbackData: create_librarian_source},
			},
			{
				{Text: LibraryInformation, CallbackData: default_menu}, //TODO сделать handler
			},
			{
				{Text: CreateSourceHowTo, CallbackData: default_menu}, //TODO сделать handler
			},
			{
				{Text: BackTo + StarMenu, CallbackData: general_start},
			},
			{
				{Text: BackTo + ReconComGroup, URL: "https://t.me/+qbEymR_JfXFhOWUy"}, //TODO пересобрать на адрес из переменной окружения
			},
		},
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        Library,
		ReplyMarkup: kb,
		ParseMode:   models.ParseModeMarkdown,
	})
}
