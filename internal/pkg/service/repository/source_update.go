package repository

import (
	"TG_Bot_Admin/internal/pkg/service/repository/dto"
	"TG_Bot_Admin/internal/pkg/service/repository/query/sources"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// UpdateSource ...
func (r *repository) UpdateSource(ctx context.Context, q sources.Update) (dto.Sources, error) {
	var res dto.Sources

	if err := Selectx(ctx, r.storage, &res, updateSourceQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, updateSourceQuery(q)): %w", err)
	}

	return res, nil
}

func updateSourceQuery(query sources.Update) sq.UpdateBuilder {
	updateQuery := sq.StatementBuilder.Update(dto.SourcesTableName).
		Prefix("--UpdateSource\n")

	for _, source := range query.Sources {

		// что дополнять
		if source.ID != 0 {
			updateQuery = updateQuery.Where(sq.Eq{dto.SourcesColumnID: source.ID})
		}
		if source.UserID != 0 {
			updateQuery = updateQuery.Where(sq.Eq{dto.SourcesColumnUserID: source.UserID})
		}

		// чем дополнять
		if source.Step != 0 {
			updateQuery = updateQuery.Set(dto.SourcesColumnStep, source.Step)
		}
		if source.Type != "" {
			updateQuery = updateQuery.Set(dto.SourcesColumnType, source.Type)
		}
		if source.NameRU != "" {
			updateQuery = updateQuery.Set(dto.SourcesColumnNameRU, source.NameRU)
		}
		if source.NameENG != "" {
			updateQuery = updateQuery.Set(dto.SourcesColumnNameENG, source.NameENG)
		}
		if source.AuthorRU != "" {
			updateQuery = updateQuery.Set(dto.SourcesColumnAuthorRU, source.AuthorRU)
		}
		if source.Year != 0 {
			updateQuery = updateQuery.Set(dto.SourcesColumnYear, source.Year)
		}
		if source.Regions != "" {
			updateQuery = updateQuery.Set(dto.SourcesColumnRegions, source.Regions)
		}
		if source.TimePeriods != "" {
			updateQuery = updateQuery.Set(dto.SourcesColumnTimePeriods, source.TimePeriods)
		}
		if source.Description != "" {
			updateQuery = updateQuery.Set(dto.SourcesColumnDescription, source.Description)
		}
		if source.FileFormat != "" {
			updateQuery = updateQuery.Set(dto.SourcesColumnFileFormat, source.FileFormat)
		}
		if source.CreatedAt != nil {
			updateQuery = updateQuery.Set(dto.SourcesColumnCreatedAt, source.CreatedAt)
		}
		if source.IsSent != 0 {
			updateQuery = updateQuery.Set(dto.SourcesColumnIsSent, source.IsSent)
		}

		updateQuery = updateQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.SourcesColumnID))
	}

	return updateQuery
}
