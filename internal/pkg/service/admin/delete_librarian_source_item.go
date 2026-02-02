package admin

import (
	"TG_Bot_Admin/internal/pkg/service/repository/query/sources"
	"context"
)

// DeleteLibrarianSourceItem ...
func (a *adminService) DeleteLibrarianSourceItem(ctx context.Context, userID int64) int64 {
	source, err := a.repository.DeleteSource(ctx, sources.Delete{
		UserIDs: []int64{userID},
	})
	if err != nil || source == nil {
		return 0
	}

	return source.Entity()[0].ID
}
