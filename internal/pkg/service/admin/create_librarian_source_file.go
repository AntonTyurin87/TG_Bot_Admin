package admin

import (
	"TG_Bot_Admin/internal/pkg/domain/entity"
	"TG_Bot_Admin/internal/pkg/service/repository/query/files"
	"TG_Bot_Admin/internal/pkg/service/repository/query/sources"
	"context"
)

func (a *adminService) CreateLibrarianSourceFile(ctx context.Context, fileData []byte, userID int64) (*entity.File, error) {
	if fileData == nil || userID == 0 {
		return nil, nil // TODO логирование ошибок
	}

	// идём в БД за источником по userID
	source, err := a.repository.SelectSource(ctx, sources.Select{
		UserIDs: []int64{userID},
	})
	if err != nil || len(source) == 0 || source.Entity()[0].Step != entity.SourceDescriptionStep {
		return nil, nil // TODO логирование ошибок
	}

	file, err := a.repository.InsertFile(ctx, files.Insert{
		Files: []*entity.File{
			{
				SourceID: source.Entity()[0].ID,
				FileData: fileData,
			},
		},
	})
	if err != nil {
		return nil, nil // TODO логирование ошибок
	}

	return file.Entity()[0], nil
}
