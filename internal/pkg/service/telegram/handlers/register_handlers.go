package telegram

import (
	"github.com/go-telegram/bot"
)

// registerHandlers регистрирует обработчики команд
func (h *Handler) RegisterHandlers(b *bot.Bot) {
	// Вызов бота и начало работы с ним
	b.RegisterHandler(bot.HandlerTypeMessageText, admin_topic_start, bot.MatchTypeExact, h.adminBotStartHandler)

	// Распределитель по меню в зависимости от прав
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, general_start, bot.MatchTypeExact, h.generalStartHandler)

	// Меню для SuperAdmin (SuperAdmin + LibrarianAdmin + SimpleUser)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, super_admin_start, bot.MatchTypeExact, h.generalStartHandler)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, super_admin_library, bot.MatchTypeExact, h.superAdminLibraryHandler)

	// Меню для LibrarianAdmin (LibrarianAdmin + SimpleUser)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, create_librarian_source, bot.MatchTypeExact, h.adminCreateLibrarianSourceHandler)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, create_librarian_book_source, bot.MatchTypeExact, h.createLibrarianBookSourceHandler)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, create_source_default, bot.MatchTypeExact, h.createLibrarianBookSourceHandler)
	// Меню для SimpleUser

	//Обработчик всех текстовых сообщений
	b.RegisterHandler(bot.HandlerTypeMessageText, text_input, bot.MatchTypePrefix, h.messageHandler)

	// Меню по умолчанию
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, default_menu, bot.MatchTypeExact, h.defaultHandler)

	//b.RegisterHandler(bot.HandlerTypeMessageText, "/menu", bot.MatchTypeExact, h.menuHandler)

	//// Регистрируем обработчики callback-ов
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "menu", bot.MatchTypeExact, h.menuCallbackHandler)
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "search", bot.MatchTypeExact, searchCallbackHandler)
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "search_item_card", bot.MatchTypeExact, searchItemCardCallbackHandler)
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "about", bot.MatchTypeExact, sectionCallbackHandler)
	//
	//// Обработчики для library
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "library", bot.MatchTypeExact, libraryCallbackHandler)
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "new_sources", bot.MatchTypeExact, h.newLibrarySourcesHandler)
	//
	//// Обработчики для work_with_sources
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "work_with_sources", bot.MatchTypeExact, h.workWithSourcesHendler)
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "add_text_source", bot.MatchTypeExact, h.addTextSourceHandler) //TODO сделать handler
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "add_text_source_book", bot.MatchTypeExact, h.addTextSourceBookHandler)
	////b.RegisterHandler(bot.HandlerTypeCallbackQueryData, btnShowText, bot.MatchTypeExact, showTextHandler)                       //TODO возможно удалить или разобраться
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "add_photo_source", bot.MatchTypeExact, sectionCallbackHandler)         //TODO сделать handler
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "add_period_to_source", bot.MatchTypeExact, sectionCallbackHandler)     //TODO сделать handler
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "add_item_group_to_source", bot.MatchTypeExact, sectionCallbackHandler) //TODO сделать handler
	//
	//// Новые обработчики для подменю поиска
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "search_period", bot.MatchTypeExact, h.searchPeriodCallbackHandler)
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "search_subject", bot.MatchTypeExact, h.searchSubmenuCallbackHandler)
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "back_to_search", bot.MatchTypeExact, searchCallbackHandler)
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "back_to_menu", bot.MatchTypeExact, h.menuCallbackHandler)
	//
	//// Обработчики для данных поиска
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "search_region", bot.MatchTypeExact, h.searchRegionCallbackHandler)
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "refresh_search_data", bot.MatchTypeExact, h.refreshSearchDataHandler)
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "use_cached_data", bot.MatchTypeExact, h.useCachedDataHandler)
	//
	//// Регистрируем обработчик для выбора региона
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "region_", bot.MatchTypePrefix, h.regionSelectedCallbackHandler)
	//
	//// Обработчик для поиска в конкретном регионе
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "search_in_region_", bot.MatchTypePrefix, h.searchInRegionHandler)
	//
	////ВНИМАНИЕ: period_group_ должен быть ЗАРЕГИСТРИРОВАН ДО period_
	//// так как period_ будет перехватывать period_group_
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "period_group_", bot.MatchTypePrefix, h.periodGroupSelectedHandler)
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "period_", bot.MatchTypePrefix, h.periodSelectedHandler)
	//
	//// Обработчики для групп периодов (после выбора региона)
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "refresh_periods_", bot.MatchTypePrefix, h.refreshPeriodsHandler)
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "search_in_group_", bot.MatchTypePrefix, h.searchInGroupHandler)
	//
	//// Обработчики для групп предметов
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "item_group_", bot.MatchTypePrefix, h.itemGroupSelectedHandler)
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "refresh_item_groups_", bot.MatchTypePrefix, h.refreshItemGroupsHandler)
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "search_items_", bot.MatchTypePrefix, h.searchItemsHandler)
	//
	//// Обработчики для поиска по всем периодам/предметам
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "search_all_periods_", bot.MatchTypePrefix, h.searchAllPeriodsHandler)
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "search_all_items_", bot.MatchTypePrefix, h.searchAllItemsHandler)
	//

}
