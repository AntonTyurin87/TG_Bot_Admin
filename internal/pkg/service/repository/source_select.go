package repository

import (
	"TG_Bot_Admin/internal/pkg/service/repository/dto"
	"TG_Bot_Admin/internal/pkg/service/repository/query/sources"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// SelectSource ...
func (r *repository) SelectSource(ctx context.Context, q sources.Select) (dto.Sources, error) {
	var res dto.Sources

	if err := Selectx(ctx, r.storage, &res, selectSourceQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, selectSourceQuery(q)): %w", err)
	}

	return res, nil
}

func selectSourceQuery(query sources.Select) sq.SelectBuilder {
	selectQuery := sq.StatementBuilder.Select(dto.SourcesColumns...).
		From(dto.SourcesTableName).
		Prefix("--SelectSource\n")

	where := sq.Eq{}

	if len(query.IDs) > 0 {
		where[dto.SourcesColumnID] = query.IDs
	}

	if len(query.UserIDs) > 0 {
		where[dto.SourcesColumnUserID] = query.UserIDs
	}

	selectQuery = selectQuery.Where(where)
	selectQuery = selectQuery.OrderBy(dto.SourcesColumnID)

	return selectQuery
}
