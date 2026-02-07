package telegram

import (
	"TG_Bot_Admin/internal/pkg/domain/entity"
	"TG_Bot_Admin/internal/pkg/domain/menu"
	"TG_Bot_Admin/internal/pkg/service/telegram/auth"
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// UserSourceStates - Переменная для хранения информации о диалоге с пользователем.
var UserSourceStates = make(map[int64]menu.KeyName)

func (h *Handler) createLibrarianBookSourceHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})

	userID := update.CallbackQuery.From.ID
	chatID := update.CallbackQuery.Message.Message.Chat.ID

	if auth.SuperAdmin != auth.GetUserCategory(userID) {
		// отправляемся в меню по умолчанию
		h.DefaultAnswerMenu(ctx, b, chatID, librarian_admin_start)
	}

	// проверка, что у данного пользователя ещё нет заготовки источника
	if h.adminService.IsAnyNotFinishedSource(ctx, userID) {
		h.continueSourceCreatingMenu(ctx, b, chatID, userID)
		return
	}

	h.createLibrarianBookSourceMenu(ctx, b, chatID, userID)
}

func (h *Handler) createLibrarianBookSourceMenu(ctx context.Context, b *bot.Bot, chatID, userID int64) {
	//создаём запись о книге.
	book, err := h.adminService.CreateLibrarianSourceItem(ctx, entity.BookSourceType, userID)
	if err != nil {
		fmt.Errorf("create librarian book source menu error: %w", err)
	}

	// получаем текст для заголовка сообщения
	messageText := h.presenter.TextMessageToCreateSource(book)

	kb := h.presenter.KeyBlockToCreateSource(book)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        messageText,
		ReplyMarkup: kb,
		ParseMode:   models.ParseModeMarkdown,
	})
}

func (h *Handler) continueSourceCreatingMenu(ctx context.Context, b *bot.Bot, chatID, userID int64) {
	source, err := h.adminService.SelectLibrarianSourceItem(ctx, userID)
	if err != nil {
		fmt.Errorf("h.adminService.SelectLibrarianSourceItem: %w", err)
	}

	if source == nil {
		return
	}

	// получаем текст для заголовка сообщения
	messageText := h.presenter.TextMessageToContinueSource(source)

	kb := h.presenter.KeyBlockToCreateSource(source)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        messageText,
		ReplyMarkup: kb,
		ParseMode:   models.ParseModeMarkdown,
	})
}
