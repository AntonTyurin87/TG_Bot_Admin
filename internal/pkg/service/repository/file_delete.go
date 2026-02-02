package repository

import (
	"TG_Bot_Admin/internal/pkg/service/repository/dto"
	"TG_Bot_Admin/internal/pkg/service/repository/query/files"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// DeleteFile ...
func (r *repository) DeleteFile(ctx context.Context, q files.Delete) (dto.Files, error) {
	var res dto.Files

	if err := Selectx(ctx, r.storage, &res, deleteFileQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, deleteFileQuery(q)): %w", err)
	}

	return res, nil
}

func deleteFileQuery(query files.Delete) sq.DeleteBuilder {
	deleteQuery := sq.StatementBuilder.Delete(dto.FilesTableName).
		Prefix("--DeleteFile\n")

	where := sq.Eq{}

	if len(query.IDs) > 0 {
		where[dto.FilesColumnID] = query.IDs
	}

	if len(query.SourceIDs) > 0 {
		where[dto.FilesColumnSourceID] = query.SourceIDs
	}

	deleteQuery = deleteQuery.Where(where)

	deleteQuery = deleteQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.FilesColumnID))

	return deleteQuery
}
