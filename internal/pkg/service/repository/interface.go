package repository

import (
	"TG_Bot_Admin/internal/pkg/service/repository/dto"
	"TG_Bot_Admin/internal/pkg/service/repository/query/sources"
	"context"
)

type Interface interface {
	InsertSource(ctx context.Context, q sources.Insert) (dto.Sources, error)
	SelectSource(ctx context.Context, q sources.Select) (dto.Sources, error)
	UpdateSource(ctx context.Context, q sources.Update) (dto.Sources, error)
	DeleteSource(ctx context.Context, q sources.Delete) (dto.Sources, error)
}
