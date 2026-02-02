package admin

import (
	"TG_Bot_Admin/internal/pkg/service/repository/query/sources"
	"context"
)

// IsAnyNotFinishedSource ...
func (a *adminService) IsAnyNotFinishedSource(ctx context.Context, userID int64) bool {
	if userID == 0 {
		return false
	}

	source, err := a.repository.SelectSource(ctx, sources.Select{
		UserIDs: []int64{userID},
	})
	if err != nil {
		return false
	}

	return len(source.Entity()) != 0
}
