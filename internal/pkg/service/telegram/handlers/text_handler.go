package telegram

import (
	"TG_Bot_Admin/internal/pkg/domain/entity"
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

	// Проверяем состояние пользователя и источника
	if h.adminService.IsAnyNotFinishedSource(ctx, userID) && h.adminService.IsStepLessThen(ctx, entity.SourceReadyToSend, userID) {
		// записываем текс в БД
		h.adminService.UpdateLibrarianSourceItem(ctx, userID, text)
		// вызываем дефолтный хендлер для сохранения источника
		h.createSourceDefaultHandler(ctx, b, update)
	}
}

// TODO передать в Librarian НЕ УДАЛЯТЬ!
//func (h *Handler) createLibrarianSourceFile(ctx context.Context, text string, userID int64) (string, error) {
//	fileURL := helpers.PrepareURLForDownload(text)
//
//	fileData, fileExtension, err := helpers.GetFileBytes(fileURL)
//	if err != nil {
//		return "", nil //TODO как-то обрабатывать эту ситуацию
//	}
//
//	return fileExtension, nil
//}
