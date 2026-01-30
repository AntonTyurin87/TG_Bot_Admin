package helpers

import "strings"

// EscapeMarkdown - функция для экранирования специальных символов Markdown
func EscapeMarkdown(text string) string {
	// Список символов, которые нужно экранировать
	chars := []string{"_", "[", "]", "(", ")", "~", "`", ">", "#", "+", "-", "=", "|", "{", "}", ".", "!"}
	for _, char := range chars {
		text = strings.ReplaceAll(text, char, "\\"+char)
	}
	return text
}
