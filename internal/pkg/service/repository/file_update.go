package repository

import (
	"TG_Bot_Admin/internal/pkg/service/repository/dto"
	"TG_Bot_Admin/internal/pkg/service/repository/query/files"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// UpdateFile ...
func (r *repository) UpdateFile(ctx context.Context, q files.Update) (dto.Files, error) {
	var res dto.Files

	if err := Selectx(ctx, r.storage, &res, updateFileQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, updateFileQuery(q)): %w", err)
	}

	return res, nil
}

func updateFileQuery(query files.Update) sq.UpdateBuilder {
	updateQuery := sq.StatementBuilder.Update(dto.FilesTableName).
		Prefix("--UpdateFile\n")

	for _, file := range query.Files {
		// что дополнять
		if file.ID != 0 {
			updateQuery = updateQuery.Where(sq.Eq{dto.FilesColumnID: file.ID})
		}
		if file.SourceID != 0 {
			updateQuery = updateQuery.Where(sq.Eq{dto.FilesColumnSourceID: file.SourceID})
		}

		// чем дополнять
		if file.FileData != nil {
			updateQuery = updateQuery.Set(dto.FilesColumnFileData, file.FileData)
		}

		updateQuery = updateQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.FilesColumnID))
	}

	return updateQuery
}
