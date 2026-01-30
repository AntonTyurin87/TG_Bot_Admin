package repository

import (
	"TG_Bot_Admin/internal/pkg/service/storage"
)

type repository struct {
	storage storage.Storage
}

func NewRepository(
	storage storage.Storage,
) *repository {
	return &repository{
		storage: storage,
	}
}
