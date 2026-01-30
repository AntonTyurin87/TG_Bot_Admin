package helpers

import "github.com/go-telegram/bot/models"

// GetUserName - вспомогательная функция для получения имени пользователя
func GetUserName(user *models.User) string {
	if user == nil {
		return "друг"
	}
	if user.FirstName != "" {
		return EscapeMarkdown(user.FirstName)
	}
	if user.Username != "" {
		return "@" + user.Username
	}
	return "друг"
}
