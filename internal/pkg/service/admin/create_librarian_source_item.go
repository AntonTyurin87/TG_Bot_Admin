package admin

import (
	"TG_Bot_Admin/internal/pkg/domain/entity"
	"TG_Bot_Admin/internal/pkg/service/repository/query/sources"
	"context"
	"fmt"
)

// CreateLibrarianSourceItem ...
func (a *adminService) CreateLibrarianSourceItem(ctx context.Context, sourceType entity.SourceType, userID int64) (*entity.Source, error) {
	err := validateLibrarianSourceItemData(sourceType, userID)
	if err != nil {
		return nil, err
	}

	source, err := a.repository.InsertSource(ctx, sources.Insert{Sources: []*entity.Source{{
		UserID: userID,
		Type:   sourceType,
		Step:   entity.CreateSourceStep,
	}}})
	if err != nil {
		return nil, err
	}

	if len(source) == 0 {
		return nil, nil
	}

	//создан точно только один источник
	return source[0].Entity(), nil
}

func validateLibrarianSourceItemData(sourceType entity.SourceType, userID int64) error {
	if sourceType == entity.UnknownSourceType {
		fmt.Errorf("unknow source type")
	}

	if userID == 0 {
		fmt.Errorf("invalid user id")
	}

	return nil
}
