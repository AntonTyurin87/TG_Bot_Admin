package admin

import (
	"TG_Bot_Admin/internal/pkg/domain/entity"
	"TG_Bot_Admin/internal/pkg/service/repository/query/sources"
	"context"
)

// UpdateLibrarianSourceItem - ошибки обрабатываем, но не возвращаем
func (a *adminService) UpdateLibrarianSourceItem(ctx context.Context, userID int64, text string) {
	// идём в БД за источником по userID
	source, err := a.repository.SelectSource(ctx, sources.Select{
		UserIDs: []int64{userID},
	})
	if err != nil || len(source) == 0 || source.Entity()[0].Step == entity.UnknownStep {
		return
	}

	nextStep := source.Entity()[0].Step.GetNextStep()

	// валидируем пришедший текст в зависимости от ожидаемого шага
	if !isValidText(text, nextStep) {
		return
	}

	// подготовка текста в зависимости от шага
	data := a.presenter.PrepareUpdateSourceData(source.Entity()[0], text, nextStep)

	// обновляем источник в базе
	_, err = a.repository.UpdateSource(ctx, data)
	if err != nil {
	}
}

// TODO сделать нормальную валидацию
func isValidText(text string, step entity.Step) bool {
	if text == "" {
		return false
	}

	return true
}
