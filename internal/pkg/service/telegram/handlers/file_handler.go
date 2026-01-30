package telegram

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Обработчик файловых сообщений
func (h *Handler) fileMessageHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	userID := update.Message.From.ID

	if update.Message == nil || update.Message.Document == nil {
		return
	}

	// Проверяем состояние пользователя
	if state, ok := UserSourceStates[userID]; ok && state == CreateSource {
		doc := update.Message.Document
		filePath, err := getFileBytes(ctx, b, doc.FileID)
		if err != nil {
			return
		}
		_, err = h.adminService.CreateLibrarianSourceFile(ctx, filePath, userID)
		if err != nil {
			return
		}

		// записываем в БД
		h.adminService.UpdateLibrarianSourceItem(ctx, userID, doc.MimeType)
		// вызываем дефолтный хендлер для сохранения источника
		h.createSourceDefaultHandler(ctx, b, update)
	}
}

// getFileBytes получает файл в виде []byte
func getFileBytes(ctx context.Context, b *bot.Bot, fileID string) ([]byte, error) {
	// Получаем информацию о файле
	file, err := b.GetFile(ctx, &bot.GetFileParams{
		FileID: fileID,
	})
	if err != nil {
		return nil, fmt.Errorf("не удалось получить информацию о файле: %w", err)
	}

	// Экранируем специальные символы в пути
	escapedPath := url.PathEscape(file.FilePath)

	// Формируем URL для скачивания
	fileURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", os.Getenv("TG_BOT_ADMIN_TOKEN"), escapedPath)

	// Скачиваем файл
	resp, err := http.Get(fileURL)
	if err != nil {
		return nil, fmt.Errorf("ошибка скачивания: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка сервера: %s", resp.Status)
	}

	// Читаем файл в []byte
	fileBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	return fileBytes, nil
}
