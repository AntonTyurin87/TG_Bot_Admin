package presenter

import (
	"TG_Bot_Admin/internal/pkg/domain/entity"
	"TG_Bot_Admin/internal/pkg/domain/menu"
	"TG_Bot_Admin/internal/pkg/domain/texts"
	"TG_Bot_Admin/internal/pkg/service/repository/query/sources"
	"TG_Bot_Admin/internal/pkg/service/telegram/helpers"
	"fmt"
	"strconv"
)

// KayNameFromCreateSource ...
func (p *presenter) KayNameFromCreateSource(source *entity.Source) string {
	return p.CollbackKeyNameBySourceStep(source.Step).String()
}

// TextMessageToCreateSource ...
func (p *presenter) TextMessageToCreateSource(source *entity.Source) string {
	title := p.CollbackKeyNameBySourceType(source.Type)
	text := p.InstructionsBySourceStep(source.Step)

	sourceState := p.SourceStateText(source)

	message := helpers.EscapeMarkdown(fmt.Sprintf("%s\n\n%s\n%s", title, sourceState, text))

	if text == texts.InstructionsAddSourceSuccess {
		return fmt.Sprint(message, menu.SaveLibrarianSource)
	}

	return message
}

// PrepareUpdateSourceData ...
func (p *presenter) PrepareUpdateSourceData(source *entity.Source, data string, nextStep entity.Step) sources.Update {
	switch nextStep {
	case entity.SourceNameRuStep:
		source.NameRU = data
	case entity.SourceNameENGStep:
		source.NameENG = data
	case entity.SourceAuthorsRUStep:
		source.AuthorRU = data
	case entity.SourceYearStep:
		source.Year, _ = strconv.ParseInt(data, 10, 64) //ошибку не обрабатываем
	case entity.SourceDescriptionStep:
		source.Description = data
	case entity.SourceLoadFileStep:
		source.FileFormat = data
	}

	source.Step = nextStep

	return sources.Update{
		Sources: []*entity.Source{
			source,
		},
	}
}

// SourceStateText ...
func (p *presenter) SourceStateText(source *entity.Source) string {
	if source == nil {
		return ""
	}

	text := fmt.Sprintf(
		"*Название:* %s\n"+
			"*Назввание ENG:* %s\n"+
			"*Автор(ы):* %s\n"+
			"*Год издания:* %d\n"+
			"*Описание:* %s\n", //TODO что-то про файл
		source.NameRU, source.NameENG, source.AuthorRU, source.Year, source.Description)

	if source.FileFormat != "" {
		return text + "*Файл источника:* загружен\n"
	}

	return text
}
