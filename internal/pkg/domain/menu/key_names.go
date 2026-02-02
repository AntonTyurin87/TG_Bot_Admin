package menu

// KeyName - названия кнопок
type KeyName string

const (
	UnknownKeyName KeyName = "Неизвестная кнопка"

	CreateSource        KeyName = "➕ Добавить источник"
	DoNotCreateSource   KeyName = "🙅 Не создавать это источник"
	SaveLibrarianSource KeyName = "💾 Сохранить источник"

	CreateBookSource     KeyName = "📘 Добавить книгу"
	CreateArticleSource  KeyName = "📒 Добавить статью"
	CreateFragmentSource KeyName = "📜 Добавить фрагмент"
	CreateGraphicSource  KeyName = "🖼 Добавить графический источник"
	CreateCardSource     KeyName = "📰 Добавить карточку на источник"

	// Работа с текстовыми источниками
	KeyNameAddSourceNameRU      KeyName = "Ввести название"
	KeyNameAddSourceNameENG     KeyName = "Ввести название \"ENG\""
	KeyNameAddSourceAuthors     KeyName = "Ввести автора"
	KeyNameAddSourceYear        KeyName = "Ввести год издания"
	KeyNameAddSourceDescription KeyName = "Ввести описание к источнику"
	KeyNameAddSourceFile        KeyName = "Послать файл источника "
	KeyNameSourceSuccess        KeyName = "Посмотреть на источник"
	KeyNameDeleteDraftSource    KeyName = "Удалить заготовку источника"
)

// String ...
func (k KeyName) String() string {
	return string(k)
}
