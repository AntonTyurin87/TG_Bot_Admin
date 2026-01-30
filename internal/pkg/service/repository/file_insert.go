package repository

import (
	"TG_Bot_Admin/internal/pkg/service/repository/dto"
	"TG_Bot_Admin/internal/pkg/service/repository/query/files"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// InsertFile ...
func (r *repository) InsertFile(ctx context.Context, q files.Insert) (dto.Files, error) {
	var res dto.Files

	if err := Selectx(ctx, r.storage, &res, insertFileQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, insertFileQuery(q)): %w", err)
	}

	return res, nil
}

func insertFileQuery(query files.Insert) sq.InsertBuilder {
	insertQuery := sq.StatementBuilder.Insert(dto.FilesTableName).
		Columns(
			dto.FilesColumnSourceID,
			dto.FilesColumnFileData,
		).
		Prefix("--InsertFile\n")

	for _, file := range query.Files {
		a := dto.FileDtoFromEntity(file)
		insertQuery = insertQuery.Values(
			a.GetSourceID(),
			a.GetFileData(),
		)
	}

	insertQuery = insertQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.FilesColumnID))

	return insertQuery
}
