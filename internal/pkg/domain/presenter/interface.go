package presenter

import (
	"TG_Bot_Admin/internal/pkg/domain/entity"
	"TG_Bot_Admin/internal/pkg/service/repository/query/sources"
)

type Interface interface {
	KayNameFromCreateSource(source *entity.Source) string
	TextMessageToCreateSource(source *entity.Source) string
	TextMessageToContinueSource(source *entity.Source) string
	PrepareUpdateSourceData(source *entity.Source, text string, nextStep entity.Step) sources.Update
}
