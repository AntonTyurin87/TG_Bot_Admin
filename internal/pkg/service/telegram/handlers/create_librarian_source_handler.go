package telegram

import (
	"TG_Bot_Admin/internal/pkg/service/telegram/auth"
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (h *Handler) adminCreateLibrarianSourceHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})

	userID := update.CallbackQuery.From.ID
	chatID := update.CallbackQuery.Message.Message.Chat.ID

	if auth.SuperAdmin != auth.GetUserCategory(userID) {
		// отправляемся в меню по умолчанию
		h.DefaultAnswerMenu(ctx, b, chatID, librarian_admin_start)
	}

	h.adminCreateLibrarianSourceMenu(ctx, b, chatID)
}

func (h *Handler) adminCreateLibrarianSourceMenu(ctx context.Context, b *bot.Bot, chatID int64) {

	//TODO Если ранее создавался источник с этим админом, то предложить вернуться к нему или удалить его и продолжить.

	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: CreateBookSource, CallbackData: create_librarian_book_source},
			},
			{
				{Text: CreateArticleSource, CallbackData: default_menu}, //TODO сделать handler
			},
			{
				{Text: CreateFragmentSource, CallbackData: default_menu}, //TODO сделать handler
			},
			{
				{Text: CreateGraphicSource, CallbackData: default_menu}, //TODO сделать handler
			},
			{
				{Text: CreateCardSource, CallbackData: default_menu}, //TODO сделать handler
			},
			{
				{Text: BackTo + Library, CallbackData: general_start},
			},
			{
				{Text: BackTo + ReconComGroup, URL: "https://t.me/+qbEymR_JfXFhOWUy"}, //TODO пересобрать на адрес из переменной окружения
			},
		},
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        CreateSource, //TODO вернуться в меню библиотеки на разделение по правам!
		ReplyMarkup: kb,
		ParseMode:   models.ParseModeMarkdown,
	})
}
