package repository

import (
	"TG_Bot_Admin/internal/pkg/service/repository/dto"
	"TG_Bot_Admin/internal/pkg/service/repository/query/files"
	"TG_Bot_Admin/internal/pkg/service/repository/query/sources"
	"context"
)

type Interface interface {
	InsertSource(ctx context.Context, q sources.Insert) (dto.Sources, error)
	SelectSource(ctx context.Context, q sources.Select) (dto.Sources, error)
	UpdateSource(ctx context.Context, q sources.Update) (dto.Sources, error)
	DeleteSource(ctx context.Context, q sources.Delete) (dto.Sources, error)

	InsertFile(ctx context.Context, q files.Insert) (dto.Files, error)
	DeleteFile(ctx context.Context, q files.Delete) (dto.Files, error)
}
