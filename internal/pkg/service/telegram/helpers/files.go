package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"
)

// GetFileBytes получает файл в виде []byte
func GetFileBytes(fileURL string) ([]byte, string, error) {
	// Скачиваем файл
	resp, err := http.Get(fileURL)
	if err != nil {
		return nil, "", fmt.Errorf("ошибка скачивания: %w", err)
	}
	defer resp.Body.Close()

	fileExtension, err := GetExtensionFromResponse(resp)
	if err != nil {
		return nil, "", fmt.Errorf("ошибка получения расширения файла: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("ошибка сервера: %s", resp.Status)
	}

	// Читаем файл в []byte
	fileBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", fmt.Errorf("ошибка чтения файла: %w", err)
	}

	//TODO подумать про статистику загрузки файла

	return fileBytes, fileExtension, nil
}

func ifGoogleURL(url string) bool {
	// Проверяем наличие ключевых слов в URL
	downloadKeywords := []string{
		"google.com", // FTP ссылки
	}

	textLower := strings.ToLower(url)

	for _, keyword := range downloadKeywords {
		if strings.Contains(textLower, keyword) {
			return true
		}
	}

	return false
}

func ifYandexURL(url string) bool {
	// Проверяем наличие ключевых слов в URL
	downloadKeywords := []string{
		"disk.yandex.ru/d/",
		"disk.yandex.ru/i/",
		"yadi.sk/d/",
		"yadi.sk/i/",
		"disk.yandex.com/d/",
		"disk.yandex.com/i/", // FTP ссылки
	}

	textLower := strings.ToLower(url)

	for _, keyword := range downloadKeywords {
		if strings.Contains(textLower, keyword) {
			return true
		}
	}

	return false
}

func prepareGoogleURL(url string) string {
	downloadKeywords := []string{"=download&"}

	for _, keyword := range downloadKeywords {
		if strings.Contains(url, keyword) {
			return url
		}
	}

	url = strings.ReplaceAll(url, "file/d/", "uc?export=download&id=")
	url = strings.TrimSuffix(url, "/view?usp=drive_link")

	return url
}

// https://cloud-api.yandex.net/v1/disk/public/resources/download?public_key=https://disk.yandex.ru/i/aIoQj6CK6l87zg"

// YandexDiskDownloadResponse - Структура для парсинга ответа Яндекс.Диска
type YandexDiskDownloadResponse struct {
	Method    string `json:"method"`
	Href      string `json:"href"`
	Templated bool   `json:"templated"`
}

func getDataFromYandexDisk(publicURL string) (string, error) {
	// Проверяем, что URL валиден
	_, err := url.Parse(publicURL)
	if err != nil {
		return "", fmt.Errorf("неверный URL: %v", err)
	}

	// Кодируем публичную ссылку для query параметра
	encodedURL := url.QueryEscape(publicURL)

	// Формируем URL для запроса к API Яндекс.Диска
	apiURL := fmt.Sprintf("https://cloud-api.yandex.net/v1/disk/public/resources/download?public_key=%s", encodedURL)

	// Создаем HTTP клиент с настройками
	client := &http.Client{
		Timeout: 30 * time.Second,
		// Можно добавить дополнительные настройки:
		// CheckRedirect: func(req *http.Request, via []*http.Request) error {
		//     return http.ErrUseLastResponse
		// },
	}

	// Создаем GET запрос
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return "", fmt.Errorf("ошибка создания запроса: %v", err)
	}

	// Добавляем заголовки (имитируем реальный браузер)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Referer", "https://disk.yandex.ru/")
	req.Header.Set("Origin", "https://disk.yandex.ru")

	// Выполняем запрос
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("ошибка выполнения HTTP запроса: %v", err)
	}
	defer resp.Body.Close()

	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения ответа: %v", err)
	}

	// Проверяем статус код
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP ошибка %d: %s", resp.StatusCode, resp.Status)
	}

	// Парсим успешный ответ
	var diskResponse YandexDiskDownloadResponse
	if err := json.Unmarshal(body, &diskResponse); err != nil {
		return "", fmt.Errorf("ошибка парсинга JSON ответа: %v\nОтвет: %s", err, string(body))
	}

	// Проверяем, что ссылка не пустая
	if diskResponse.Href == "" {
		return "", fmt.Errorf("пустая ссылка в ответе Яндекс.Диска")
	}

	// Возвращаем href
	return diskResponse.Href, nil
}

func PrepareURLForDownload(url string) string {
	// проверка и подготовка гугловской строки для скачивания
	if ifGoogleURL(url) {
		return prepareGoogleURL(url)
	}
	// проверка и подготовка яндексовой строки для скачивания
	if ifYandexURL(url) {
		fileURLYandex, err := getDataFromYandexDisk(url)
		if err != nil {
			return url //TODO подумать про обработку ошибок
		}
		return fileURLYandex
	}

	return url
}

// GetExtensionFromResponse ...
func GetExtensionFromResponse(resp *http.Response) (string, error) {
	// 1. Из заголовка Content-Disposition
	contentDisposition := resp.Header.Get("Content-Disposition")
	if contentDisposition != "" {
		// Пример: "attachment; filename=\"document.pdf\""
		// Пример: "inline; filename=report.docx"

		// Парсим значение filename
		_, params, err := mime.ParseMediaType(contentDisposition)
		if err == nil {
			if filename, ok := params["filename"]; ok && filename != "" {
				ext := filepath.Ext(filename)
				if ext != "" {
					return strings.ToLower(ext), nil
				}
			}
		}
	}

	// 2. Из заголовка Content-Type
	contentType := resp.Header.Get("Content-Type")
	if contentType != "" {
		// Убираем параметры (например: "text/html; charset=utf-8")
		if idx := strings.Index(contentType, ";"); idx != -1 {
			contentType = contentType[:idx]
		}
		contentType = strings.TrimSpace(contentType)

		// Преобразуем MIME type в расширение
		ext, err := mimeToExtension(contentType)
		if err == nil && ext != "" {
			return ext, nil
		}
	}

	// 3. Из URL (последний сегмент пути)
	if resp.Request != nil && resp.Request.URL != nil {
		urlPath := resp.Request.URL.Path
		if urlPath != "" {
			ext := filepath.Ext(urlPath)
			if ext != "" {
				return strings.ToLower(ext), nil
			}
		}
	}

	return "", fmt.Errorf("не удалось определить расширение файла")
}

// Преобразование MIME type в расширение
func mimeToExtension(mimeType string) (string, error) {
	mimeToExt := map[string]string{
		// Документы
		"application/pdf":    ".pdf",
		"application/msword": ".doc",
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document": ".docx",
		"application/vnd.ms-excel": ".xls",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":         ".xlsx",
		"application/vnd.ms-powerpoint":                                             ".ppt",
		"application/vnd.openxmlformats-officedocument.presentationml.presentation": ".pptx",
		"text/plain":      ".txt",
		"text/csv":        ".csv",
		"application/rtf": ".rtf",

		// Архивы
		"application/zip":              ".zip",
		"application/x-rar-compressed": ".rar",
		"application/x-7z-compressed":  ".7z",
		"application/x-tar":            ".tar",
		"application/gzip":             ".gz",
		"application/x-bzip2":          ".bz2",

		// Изображения
		"image/jpeg":    ".jpg",
		"image/jpg":     ".jpg",
		"image/png":     ".png",
		"image/gif":     ".gif",
		"image/bmp":     ".bmp",
		"image/webp":    ".webp",
		"image/svg+xml": ".svg",
		"image/tiff":    ".tiff",

		// Аудио
		"audio/mpeg":     ".mp3",
		"audio/mp3":      ".mp3",
		"audio/wav":      ".wav",
		"audio/x-wav":    ".wav",
		"audio/flac":     ".flac",
		"audio/aac":      ".aac",
		"audio/ogg":      ".ogg",
		"audio/x-ms-wma": ".wma",

		// Видео
		"video/mp4":        ".mp4",
		"video/x-m4v":      ".m4v",
		"video/avi":        ".avi",
		"video/x-msvideo":  ".avi",
		"video/x-matroska": ".mkv",
		"video/quicktime":  ".mov",
		"video/x-ms-wmv":   ".wmv",
		"video/webm":       ".webm",
		"video/flv":        ".flv",

		// Прочие
		"application/json":         ".json",
		"application/xml":          ".xml",
		"text/html":                ".html",
		"text/css":                 ".css",
		"application/javascript":   ".js",
		"application/octet-stream": ".bin", // Общий бинарный файл
	}

	// Приводим к нижнему регистру для сравнения
	mimeType = strings.ToLower(mimeType)

	// Ищем точное совпадение
	if ext, ok := mimeToExt[mimeType]; ok {
		return ext, nil
	}

	// Ищем частичное совпадение
	for mimePattern, ext := range mimeToExt {
		if strings.Contains(mimeType, mimePattern) {
			return ext, nil
		}
	}

	return "", fmt.Errorf("неизвестный MIME type: %s", mimeType)
}
