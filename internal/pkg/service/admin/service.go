package admin

import (
	"TG_Bot_Admin/internal/pkg/domain/presenter"
	"TG_Bot_Admin/internal/pkg/service/librarian"
	"TG_Bot_Admin/internal/pkg/service/repository"
)

type adminService struct {
	presenter  presenter.Interface
	librarian  librarian.Interface
	repository repository.Interface
}

// NewAdminService ...
func NewAdminService(
	presenter presenter.Interface,
	librarian librarian.Interface,
	repository repository.Interface,
) *adminService {
	return &adminService{
		presenter:  presenter,
		librarian:  librarian,
		repository: repository,
	}
}
