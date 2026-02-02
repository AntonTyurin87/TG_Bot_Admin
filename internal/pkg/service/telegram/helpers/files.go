package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// GetFileBytes получает файл в виде []byte
func GetFileBytes(fileURL string) ([]byte, error) {
	// проверка и подготовка гугловской строки для скачивания
	if ifGoogleURL(fileURL) {
		fileURL = prepareGoogleURL(fileURL)
	}

	if ifYandexURL(fileURL) {
		fileURL, _ = getDataFromYandexDisk(fileURL) //TODO подумать про обработку ошибок
	}

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

	//TODO подумать про статистику загрузки файла

	return fileBytes, nil
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
