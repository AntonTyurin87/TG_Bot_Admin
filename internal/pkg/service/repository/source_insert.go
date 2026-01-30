package repository

import (
	"TG_Bot_Admin/internal/pkg/service/repository/dto"
	"TG_Bot_Admin/internal/pkg/service/repository/query/sources"
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

// InsertSource ...
func (r *repository) InsertSource(ctx context.Context, q sources.Insert) (dto.Sources, error) {
	var res dto.Sources

	if err := Selectx(ctx, r.storage, &res, insertSourceQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, insertSourceQuery(q)): %w", err)
	}

	return res, nil
}

func insertSourceQuery(query sources.Insert) sq.InsertBuilder {
	insertQuery := sq.StatementBuilder.Insert(dto.SourcesTableName).
		Columns(
			dto.SourcesColumnUserID,
			dto.SourcesColumnStep,
			dto.SourcesColumnType,
			dto.SourcesColumnNameRU,
			dto.SourcesColumnNameENG,
			dto.SourcesColumnAuthorRU,
			dto.SourcesColumnYear,
			dto.SourcesColumnRegions,
			dto.SourcesColumnTimePeriods,
			dto.SourcesColumnDescription,
			dto.SourcesColumnFileFormat,
			dto.SourcesColumnCreatedAt,
			dto.SourcesColumnIsSent,
		).
		Prefix("--InsertSource\n")

	for _, source := range query.Sources {
		a := dto.SourceDtoFromEntity(source)
		insertQuery = insertQuery.Values(
			a.GetUserID(),
			a.GetStep(),
			a.GetType(),
			a.GetNameRU(),
			a.GetNameENG(),
			a.GetAuthorRU(),
			a.GetYear(),
			a.GetRegions(),
			a.GetTimePeriods(),
			a.GetDescription(),
			a.GetFileFormat(),
			a.GetCreatedAt(),
			a.GetIsSent(),
		)
	}

	insertQuery = insertQuery.Suffix(fmt.Sprintf("RETURNING %s", strings.Join(dto.SourcesColumns, ", ")))

	return insertQuery
}
