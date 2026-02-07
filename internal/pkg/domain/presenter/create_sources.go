package presenter

import (
	"TG_Bot_Admin/internal/pkg/domain/entity"
	"TG_Bot_Admin/internal/pkg/domain/menu"
	"TG_Bot_Admin/internal/pkg/domain/texts"
	"TG_Bot_Admin/internal/pkg/service/repository/query/sources"
	"TG_Bot_Admin/internal/pkg/service/telegram/helpers"
	"fmt"
	"strconv"
	"time"

	"github.com/go-telegram/bot/models"
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
	if text == texts.InstructionsSourceSendToSave {
		message = helpers.EscapeMarkdown(fmt.Sprintf("%s\n\n%s", title, texts.InstructionsSourceSendToSave.String()))
	}

	return message
}

// TextMessageToContinueSource ...
func (p *presenter) TextMessageToContinueSource(source *entity.Source) string {
	title := p.CollbackKeyNameBySourceType(source.Type)
	text := p.InstructionsBySourceStep(source.Step)

	sourceState := p.SourceStateText(source)

	message := helpers.EscapeMarkdown(fmt.Sprintf("%s\n\n%s\n\n%s\n%s", title, texts.InstructionsContinueCreatingSource, sourceState, text))

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
	case entity.SourceDownloadURLStep:
		source.DownloadURL = helpers.PrepareURLForDownload(data)
	case entity.SourceReadyToSend:
		source.CreatedAt = time.Now().Format(time.RFC3339)
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

	return text
}

// KeyBlockToCreateSource ...
func (p *presenter) KeyBlockToCreateSource(source *entity.Source) models.InlineKeyboardMarkup {
	var kb models.InlineKeyboardMarkup

	backToLibrary := []models.InlineKeyboardButton{{Text: menu.BackTo + menu.Library.String(), CallbackData: menu.General_start}}
	backToReconCom := []models.InlineKeyboardButton{{Text: menu.BackTo + menu.ReconComGroup, URL: texts.KeyURLReconComGroupURL.String()}}

	switch source.Step {
	case
		entity.CreateSourceStep,
		entity.SourceNameRuStep,
		entity.SourceNameENGStep,
		entity.SourceAuthorsRUStep,
		entity.SourceYearStep,
		entity.SourceDescriptionStep:
		kb.InlineKeyboard = append(kb.InlineKeyboard, []models.InlineKeyboardButton{{Text: menu.DoNotCreateSource.String(), CallbackData: menu.Delete_source_default}})
	case entity.SourceDownloadURLStep:
		kb.InlineKeyboard = append(kb.InlineKeyboard, []models.InlineKeyboardButton{{Text: menu.SaveLibrarianSource.String(), CallbackData: menu.Send_source_to_save}})
		kb.InlineKeyboard = append(kb.InlineKeyboard, []models.InlineKeyboardButton{{Text: menu.DoNotCreateSource.String(), CallbackData: menu.Delete_source_default}})
	default:
	}

	kb.InlineKeyboard = append(kb.InlineKeyboard, backToLibrary)
	kb.InlineKeyboard = append(kb.InlineKeyboard, backToReconCom)

	return kb
}
