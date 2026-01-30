package repository

import (
	"TG_Bot_Admin/internal/pkg/service/repository/dto"
	"TG_Bot_Admin/internal/pkg/service/repository/query/files"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// SelectFile ...
func (r *repository) SelectFile(ctx context.Context, q files.Select) (dto.Files, error) {
	var res dto.Files

	if err := Selectx(ctx, r.storage, &res, selectFilesQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, selectFilesQuery(q)): %w", err)
	}

	return res, nil
}

func selectFilesQuery(query files.Select) sq.SelectBuilder {
	selectQuery := sq.StatementBuilder.Select(dto.FilesColumns...).
		From(dto.FilesTableName).
		Prefix("--SelectFile\n")

	where := sq.Eq{}

	if len(query.IDs) > 0 {
		where[dto.FilesColumnSourceID] = query.IDs
	}

	selectQuery = selectQuery.Where(where)
	selectQuery = selectQuery.OrderBy(dto.FilesColumnID)

	return selectQuery
}
