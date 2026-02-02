package telegram

import (
	"TG_Bot_Admin/internal/pkg/domain/entity"
	"TG_Bot_Admin/internal/pkg/service/telegram/helpers"
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
	if h.adminService.IsAnyNotFinishedSource(ctx, userID) {
		// распознание ссылки для скачивания файла и готовности источника к приёму файла
		if helpers.IsDownloadLink(ctx, text) || h.adminService.IsNowStep(ctx, entity.SourceDescriptionStep, userID) { //TODO проверка на определённый шаг
			fileData, err := helpers.GetFileBytes(text)
			if err != nil {
				return //TODO как-то обрабатывать эту ситуацию
			}

			_, err = h.adminService.CreateLibrarianSourceFile(ctx, fileData, userID)
			if err != nil {
				return //TODO как-то обрабатывать эту ситуацию
			}

			text = "FileType?" //TODO передавать тип файла  (предварительно где-то его взять)
		}

		// записываем текс в БД
		h.adminService.UpdateLibrarianSourceItem(ctx, userID, text)
		// вызываем дефолтный хендлер для сохранения источника
		h.createSourceDefaultHandler(ctx, b, update)
	}
}
