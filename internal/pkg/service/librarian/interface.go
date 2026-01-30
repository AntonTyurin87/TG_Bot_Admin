package librarian

import (
	"TG_Bot_Admin/internal/pkg/domain/entity"
	"context"
)

type Interface interface {
	GetAllRegions(ctx context.Context) ([]*entity.Region, error)
}
