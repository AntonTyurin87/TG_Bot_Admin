package repository

import (
	"TG_Bot_Admin/internal/pkg/service/repository/dto"
	"TG_Bot_Admin/internal/pkg/service/repository/query/sources"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// DeleteSource ...
func (r *repository) DeleteSource(ctx context.Context, q sources.Delete) (dto.Sources, error) {
	var res dto.Sources

	if err := Selectx(ctx, r.storage, &res, deleteSourceQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, deleteSourceQuery(q)): %w", err)
	}

	return res, nil
}

func deleteSourceQuery(query sources.Delete) sq.DeleteBuilder {
	deleteQuery := sq.StatementBuilder.Delete(dto.SourcesTableName).
		Prefix("--DeleteSource\n")

	where := sq.Eq{}

	if len(query.IDs) > 0 {
		where[dto.SourcesColumnID] = query.IDs
	}
	if len(query.UserIDs) > 0 {
		where[dto.SourcesColumnUserID] = query.UserIDs
	}

	deleteQuery = deleteQuery.Where(where)

	deleteQuery = deleteQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.SourcesColumnID))

	return deleteQuery
}
