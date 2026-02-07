package admin

import (
	"TG_Bot_Admin/internal/pkg/domain/entity"
	"TG_Bot_Admin/internal/pkg/service/repository/query/sources"
	"context"
)

// IsStepLessThen ...
func (a *adminService) IsStepLessThen(ctx context.Context, step entity.Step, userID int64) bool {
	if userID == 0 {
		return false
	}

	source, err := a.repository.SelectSource(ctx, sources.Select{
		UserIDs: []int64{userID},
	})
	if err != nil {
		return false
	}

	if source.Entity()[0].Step >= step {
		return false
	}

	return true
}
