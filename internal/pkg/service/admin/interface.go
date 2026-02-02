package admin

import (
	"TG_Bot_Admin/internal/pkg/domain/entity"
	"context"
)

type Interface interface {
	CreateLibrarianSourceItem(ctx context.Context, sourceType entity.SourceType, userID int64) (*entity.Source, error)
	UpdateLibrarianSourceItem(ctx context.Context, userID int64, text string)
	SelectLibrarianSourceItem(ctx context.Context, userID int64) (*entity.Source, error)
	DeleteLibrarianSourceItem(ctx context.Context, userID int64) int64

	CreateLibrarianSourceFile(ctx context.Context, fileData []byte, userID int64) (*entity.File, error)
	DeleteLibrarianSourceFile(ctx context.Context, sourceID int64) int64

	IsAnyNotFinishedSource(ctx context.Context, userID int64) bool
	IsNowStep(ctx context.Context, step entity.Step, userID int64) bool
}
