package admin

import (
	"TG_Bot_Admin/internal/pkg/service/repository/query/files"
	"context"
)

// DeleteLibrarianSourceFile ...
func (a *adminService) DeleteLibrarianSourceFile(ctx context.Context, sourceID int64) int64 {
	file, err := a.repository.DeleteFile(ctx, files.Delete{
		SourceIDs: []int64{sourceID},
	})
	if err != nil || file == nil {
		return 0
	}

	return file.Entity()[0].ID
}
