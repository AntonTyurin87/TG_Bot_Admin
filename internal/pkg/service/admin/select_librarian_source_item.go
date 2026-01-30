package admin

import (
	"TG_Bot_Admin/internal/pkg/domain/entity"
	"TG_Bot_Admin/internal/pkg/service/repository/query/sources"
	"context"
)

// SelectLibrarianSourceItem ...
func (a *adminService) SelectLibrarianSourceItem(ctx context.Context, userID int64) (*entity.Source, error) {
	if userID == 0 {
		return nil, nil
	}

	source, err := a.repository.SelectSource(ctx, sources.Select{
		UserIDs: []int64{userID},
	})
	if err != nil {
		return nil, err
	}

	if len(source) == 0 {
		return nil, nil
	}

	//создан точно только один источник
	return source[0].Entity(), nil
}
