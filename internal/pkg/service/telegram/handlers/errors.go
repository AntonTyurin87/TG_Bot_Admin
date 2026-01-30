package telegram

//
//import (
//	"context"
//	"log"
//	"strings"
//
//	"github.com/go-telegram/bot"
//	"github.com/go-telegram/bot/models"
//)
//
//// Обработка ошибок при загрузке данных поиска
//func (h *Handler) handleSearchDataLoadError(ctx context.Context, b *bot.Bot, chatID int64, err error) {
//	errorMsg := "❌ Ошибка при загрузке данных для поиска"
//
//	if strings.Contains(err.Error(), "name resolver error") {
//		errorMsg += "\n\n🔌 Сервис поиска временно недоступен"
//	} else if strings.Contains(err.Error(), "failed to load regions") {
//		errorMsg += "\n\n⚠️ Не удалось загрузить список регионов"
//	}
//
//	// Проверяем, есть ли данные в кэше
//	cacheMutex.RLock()
//	hasCachedRegions := len(regionsCache) > 0
//	cacheMutex.RUnlock()
//
//	var keyboard [][]models.InlineKeyboardButton
//
//	if hasCachedRegions {
//		keyboard = [][]models.InlineKeyboardButton{
//			{
//				{Text: "📊 Использовать кэшированные данные", CallbackData: "use_cached_data"},
//				{Text: "🔄 Повторить попытку", CallbackData: "search_region"},
//			},
//			{
//				{Text: "🔙 Назад к поиску", CallbackData: "search"},
//				{Text: "🏠 В главное меню", CallbackData: "menu"},
//			},
//		}
//	} else {
//		keyboard = [][]models.InlineKeyboardButton{
//			{
//				{Text: "🔄 Повторить попытку", CallbackData: "search_region"},
//			},
//			{
//				{Text: "🔙 Назад к поиску", CallbackData: "search"},
//				{Text: "🏠 В главное меню", CallbackData: "menu"},
//			},
//		}
//	}
//
//	_, sendErr := b.SendMessage(ctx, &bot.SendMessageParams{
//		ChatID: chatID,
//		Text:   errorMsg,
//		ReplyMarkup: &models.InlineKeyboardMarkup{
//			InlineKeyboard: keyboard,
//		},
//	})
//
//	if sendErr != nil {
//		log.Printf("Error sending error message: %v", sendErr)
//	}
//}
