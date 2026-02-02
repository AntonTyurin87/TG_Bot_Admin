package helpers

import (
	"context"
	"strings"
)

// IsDownloadLink - Функция для проверки, является ли текст ссылкой для скачивания
func IsDownloadLink(ctx context.Context, text string) bool {
	// Проверяем, что текст не пустой
	if text == "" {
		return false
	}

	// Проверяем, что это вообще похоже на URL
	if !strings.Contains(text, "http://") && !strings.Contains(text, "https://") {
		return false
	}

	// Дополнительные проверки для ссылок скачивания
	// Проверяем расширения файлов в URL
	downloadExtensions := []string{ //TODO добавить типы файлов
		".pdf", ".zip", ".rar", ".7z", ".tar", ".gz", ".xz",
		".mp4", ".avi", ".mkv", ".mov", ".wmv",
		".mp3", ".wav", ".flac", ".aac",
		".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp",
		".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx",
		".txt", ".csv", ".json", ".xml", ".html",
	}

	textLower := strings.ToLower(text)

	// Проверяем наличие расширений файлов в URL
	for _, ext := range downloadExtensions {
		if strings.Contains(textLower, ext) {
			return true
		}
	}

	// Проверяем наличие ключевых слов в URL
	downloadKeywords := []string{
		"disk.yandex.ru",
		"/download", "/file", "/attachment", "/getfile",
		"download.php", "file.php", "attachment.php",
		"?download=", "&download=", "=download&", "?file=", "&file=",
		"download=true", "download=1",
		"/downloads/", "/files/", "/attachments/",
		"blob:", "blob:http", // Для blob ссылок
		"magnet:", // Магнет ссылки
		"ftp://",  // FTP ссылки
	}

	for _, keyword := range downloadKeywords {
		if strings.Contains(textLower, keyword) {
			return true
		}
	}

	// Проверяем Content-Disposition заголовки (если они есть в ссылке)
	if strings.Contains(textLower, "content-disposition=attachment") ||
		strings.Contains(textLower, "content-disposition=inline") {
		return true
	}

	return false
}
