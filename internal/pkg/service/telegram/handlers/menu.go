package telegram

// хендлеры для кнопок
const (
	admin_topic_start     = "/admin_topic_start"
	general_start         = "/general_start"
	super_admin_start     = "/super_admin_start"
	librarian_admin_start = "/librarian_admin_start"
	simple_user_start     = "/simple_user_start"

	super_admin_library     = "/super_admin_library"
	librarian_admin_library = "/super_admin_library"
	simple_user_library     = "/simple_user_library"

	create_librarian_source      = "/create_librarian_source"
	create_librarian_book_source = "/create_librarian_book_source"
	create_source_default        = "/create_source_default"
	delete_source_default        = "/delete_source_default"

	work_with_sources = "/work_with_sources"

	text_input = ""
	file_input = "./downloads"

	default_menu = "/default_menu"
)

// Кнопки меню
const (
	StarMenu = "Бот-Aдминистратор"

	SuperAdminStartMenu   = "Меню Супер Администратора"
	LibraryAdminStartMenu = "Меню Библиотекаря"
	SimpleUserStartMenu   = "Меню пользователя"

	Library            = "📚 Библиотека"
	LibraryInformation = "📚 Информация о библиотеке"

	CreateSourceHowTo    = "❓ Как работать с разделом"
	CreateSource         = "➕ Добавить источник"
	CreateBookSource     = "📘 Добавить книгу"
	CreateArticleSource  = "📒 Добавить статью"
	CreateFragmentSource = "📜 Добавить фрагмент"
	CreateGraphicSource  = "🖼 Добавить графический источник"
	CreateCardSource     = "📰 Добавить карточку на источник"

	BackTo = "🔙 Вернуться в "

	ReconComGroup = "группу Recom_Com"

	Empty = "Пока пустой раздел"
)
