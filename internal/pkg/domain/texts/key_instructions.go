package texts

type KeyInstruction string

// Текстовые заготовки для шагов создания текстовых источников
const (
	KeyInstructionUnknown    KeyInstruction = "К этой кнопке нет инструкции"
	KeyInstructionPushButton KeyInstruction = "Далее нажмите на кнопку и"

	KeyInstructionDeleteDraftSource KeyInstruction = "Если заготовка источника не будет заполняться далее или не верна, то можете её удалить."

	KeyURLReconComGroupURL KeyInstruction = "https://t.me/+qbEymR_JfXFhOWUy" //TODO пересобрать на адрес из переменной окружения
)

func (k KeyInstruction) String() string {
	return string(k)
}

type Instructions string

const (
	InstructionsUnknown Instructions = "InstructionsUnknown"
	InstructionsAndSend Instructions = " и нажмите *\"Send\"*"

	InstructionsAddSourceNameRU      Instructions = "Введите полное название источника на русском языке с большой буквы.\n(обязательное поле)."
	InstructionsAddSourceNameENG     Instructions = "Укажите полное название источника на английском языке с большой буквы или введите знак \"-\"."
	InstructionsAddSourceAuthors     Instructions = "Укажите фамилию и инициалы авторов, через запятую.\n(обязательное поле)."
	InstructionsAddSourceYear        Instructions = "Укажите год издания.\n(обязательное поле)."
	InstructionsAddSourceDescription Instructions = "Укажите описание источника. Не более 500 знаков."
	InstructionsAddSourceFile        Instructions = "Скопируйте файл источника и нажмите *\"Send\"*."
	InstructionsAddSourceSuccess     Instructions = "Источник успешно создан. Для загрузки в библиотеку нажмите кнопку - "
)
